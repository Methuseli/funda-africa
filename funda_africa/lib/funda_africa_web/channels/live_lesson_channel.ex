defmodule FundaAfricaWeb.LiveLessonChannel do
  use Phoenix.Channel
  alias FundaAfricaWeb.Presence
  alias FundaAfrica.LiveLessons.LiveLessonOps
  alias FundaAfrica.LiveLessons.LessonParticipantOps

  def join("live_lesson:" <> lesson_id, _params, socket) do
    user_id = socket.assigns.user_id

    # Optionally check if user is allowed in lesson
    if LiveLessonOps.lesson_active?(lesson_id) do
      # Track presence
      Presence.track(socket, user_id, %{joined_at: DateTime.utc_now()})

      # Persist participant in DB
      unless LessonParticipantOps.participant_exists?(lesson_id, user_id) do
        attrs = %{
          lesson_id: lesson_id,
          user: user_id
        }

        LessonParticipantOps.create_lesson_participant(attrs)
      end

      send(self(), :after_join)
      {:ok, socket}
    else
      {:error, %{reason: "lesson_not_active"}}
    end
  end

  # After join, broadcast presence to all participants
  def handle_info(:after_join, socket) do
    _ = socket.topic |> String.replace_prefix("live_lesson:", "")
    user_id = socket.assigns.user_id

    broadcast!(socket, "user:joined", %{user_id: user_id})
    {:noreply, socket}
  end

  # Signaling for WebRTC
  def handle_in("webrtc:offer", %{"to" => peer_id, "offer" => offer}, socket) do
    push_to_peer(peer_id, "webrtc:offer", %{
      from: socket.assigns.user_id,
      offer: offer
    })

    {:noreply, socket}
  end

  def handle_in("webrtc:answer", %{"to" => peer_id, "answer" => answer}, socket) do
    push_to_peer(peer_id, "webrtc:answer", %{
      from: socket.assigns.user_id,
      answer: answer
    })

    {:noreply, socket}
  end

  def handle_in("webrtc:candidate", %{"to" => peer_id, "candidate" => candidate}, socket) do
    push_to_peer(peer_id, "webrtc:candidate", %{
      from: socket.assigns.user_id,
      candidate: candidate
    })

    {:noreply, socket}
  end

  def handle_in("lesson:raise_hand", _params, socket) do
    broadcast!(socket, "lesson:raise_hand", %{user_id: socket.assigns.user_id})
    {:noreply, socket}
  end

  def handle_in("lesson:message", %{"content" => content}, socket) do
    lesson_id = socket.assigns.lesson_id
    user_id = socket.assigns.user_id

    # Ensure conversation exists
    conversation =
      FundaAfrica.Repo.get_by(FundaAfrica.Chat.Conversation, type: "lesson", id: lesson_id)

    # Save the message
    {:ok, msg} =
      FundaAfrica.Chat.MessageOps.create_message(%{
        conversation_id: conversation.id,
        sender: user_id,
        content: content,
        content_type: "text"
      })

    # Broadcast to lesson participants
    broadcast!(socket, "lesson:message", %{
      message_id: msg.id,
      user_id: user_id,
      content: content
    })

    {:noreply, socket}
  end

  # Chat messages during the lesson can also be transcribed if desired
  def handle_in("lesson:audio_chunk", %{"user_id" => user_id, "audio_base64" => audio}, socket) do
    # send audio to AI transcription service asynchronously
    Task.start(fn ->
      # Example: call your AI service via HTTP/gRPC
      # {:ok, transcript_text} = MyApp.AI.TranscriptionService.transcribe(audio)
      transcript_text = "Simulated transcript for audio chunk #{audio}"

      # Save transcript
      FundaAfrica.LiveLessons.LessonTranscriptOps.create_transcript(%{
        lesson_id: socket.assigns.lesson_id,
        user: user_id,
        content: transcript_text
      })

      # Broadcast transcript to lesson participants
      broadcast!(socket, "lesson:transcript", %{
        user_id: user_id,
        content: transcript_text
      })
    end)

    {:noreply, socket}
  end

  def handle_in(
        "lesson:file",
        %{"file_url" => file_url, "file_name" => file_name, "file_size" => file_size},
        socket
      ) do
    lesson_id = socket.assigns.lesson_id
    user_id = socket.assigns.user_id

    # Ensure conversation exists
    conversation =
      FundaAfrica.  Repo.get_by(FundaAfrica.Chat.Conversation, type: "lesson", title: lesson_id)

    # Persist message as a file type
    {:ok, msg} =
      FundaAfrica.Chat.MessageOps.create_message(%{
        conversation_id: conversation.id,
        sender: user_id,
        # we store URL, not the raw file
        content: file_url,
        content_type: "file",
        file_name: file_name,
        file_size: file_size,
        file_url: file_url
      })

    # Broadcast to lesson participants
    broadcast!(socket, "lesson:file", %{
      message_id: msg.id,
      user_id: user_id,
      file_name: file_name,
      file_size: file_size,
      file_url: file_url
    })

    {:noreply, socket}
  end

  def terminate(_reason, socket) do
    user_id = socket.assigns.user_id
    _ = socket.topic |> String.replace_prefix("live_lesson:", "")

    # Remove from presence
    Presence.untrack(socket, user_id)

    broadcast!(socket, "user:left", %{user_id: user_id})
    :ok
  end

  # Helper: push to peer (using presence or a topic-per-user)
  defp push_to_peer(peer_id, event, payload) do
    FundaAfricaWeb.Endpoint.broadcast("live_lesson:peer:#{peer_id}", event, payload)
  end
end
