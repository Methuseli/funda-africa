defmodule FundaAfricaWeb.LiveLessonController do
  use FundaAfricaWeb, :controller
  alias FundaAfrica.LiveLessons.LiveLessonOps
  alias FundaAfrica.LiveLessons.LessonParticipantOps
  alias FundaAfrica.LiveLessons.LessonTranscriptOps
  alias FundaAfrica.Chat.MessageOps
  alias FundaAfrica.LiveLessons.InviteOps

  action_fallback(FundaAfricaWeb.FallbackController)

  def create(conn, %{"title" => title, "host" => host, "start_time" => start_time} = params) do
    attrs = %{
      host: host,
      title: title,
      description: Map.get(params, "description"),
      start_time: start_time,
      end_time: Map.get(params, "end_time"),
      status: "scheduled"
    }

    case LiveLessonOps.create_live_lesson(attrs) do
      {:ok, lesson} ->
        # generate invite code (uuid or random string)
        attrs = %{
          lesson_id: lesson.id,
          is_active: true,
          metadata: %{}
        }
        invite = InviteOps.create_invite(attrs)
        json(conn, %{lesson: lesson, invite: invite})

      {:error, changeset} ->
        conn |> put_status(:unprocessable_entity) |> json(%{errors: changeset})
    end
  end

  def list_by_host(conn, params) do
    host_id = params["host_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()
    lessons = LiveLessonOps.list_by_host(host_id, page, page_size)
    json(conn, lessons)
  end

  def list_by_participant(conn, params) do
    user_id = params["user_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()
    lessons = LiveLessonOps.list_participant_lessons(user_id, page, page_size)
    json(conn, lessons)
  end

  def update(conn, %{"id" => lesson_id} = params) do
    lesson = LiveLessonOps.get_live_lesson!(lesson_id)

    attrs = %{
      title: Map.get(params, "title"),
      description: Map.get(params, "description"),
      start_time: Map.get(params, "start_time"),
      end_time: Map.get(params, "end_time"),
      status: Map.get(params, "status")
    }

    case LiveLessonOps.update_live_lesson(lesson, attrs) do
      {:ok, updated_lesson} ->
        json(conn, %{lesson: updated_lesson})

      {:error, changeset} ->
        conn |> put_status(:unprocessable_entity) |> json(%{errors: changeset})
    end
  end

  def show(conn, %{"id" => lesson_id}) do
    lesson = LiveLessonOps.get_live_lesson!(lesson_id)
    lesson_transcripts = LessonTranscriptOps.list_by_lesson(lesson_id)
    lesson_participants = LessonParticipantOps.list_by_lesson(lesson_id)

    messages = MessageOps.list_by_conversation(lesson_id) # if you link lesson ↔ conversation

    json(conn, %{
      lesson: lesson,
      transcripts: lesson_transcripts,
      participants: lesson_participants,
      messages: messages
    })
  end

end
