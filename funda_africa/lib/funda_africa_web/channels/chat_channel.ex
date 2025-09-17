defmodule FundaAfricaWeb.ChatChannel do
  use Phoenix.Channel
  alias FundaAfricaWeb.Presence
  alias FundaAfrica.Chat.{Memberships, MessageOps, MessageReceiptOps}

  # --- Joining topics ---
  def join("chat:lobby", _params, socket) do
    send(self(), :after_join)
    {:ok, socket}
  end

  def join("chat:org:" <> _org_id, _params, socket), do: {:ok, socket}
  def join("chat:lesson:" <> _lesson_id, _params, socket), do: {:ok, socket}
  def join("chat:private:" <> _pair, _params, socket), do: {:ok, socket}
  def join(_, _, _), do: {:error, %{reason: "unauthorized"}}

  # --- Presence tracking ---
  def handle_info(:after_join, socket) do
    user_id = socket.assigns.user_id

    Presence.track(socket, user_id, %{online_at: DateTime.utc_now()})

    for topic <- Memberships.topics_for_user(user_id) do
      FundaAfricaWeb.Endpoint.broadcast(topic, "user:online", %{user_id: user_id})
    end

    {:noreply, socket}
  end

  def terminate(_reason, socket) do
    user_id = socket.assigns.user_id

    case Presence.list(socket) |> Map.get(user_id) do
      nil ->
        for topic <- Memberships.topics_for_user(user_id) do
          FundaAfricaWeb.Endpoint.broadcast(topic, "user:offline", %{user_id: user_id})
        end

      _ ->
        :ok
    end
  end

  def handle_in(
        "message:new",
        %{
          "content" => content,
          "recipients" => recips,
          "conversation_id" => conv_id,
          "content_type" => content_type
        } = payload,
        socket
      ) do
    sender_id = socket.assigns.user_id

    # Extract optional metadata for files/videos
    file_name = Map.get(payload, "file_name")
    file_size = Map.get(payload, "file_size")
    file_url = Map.get(payload, "file_url")

    {:ok, message} =
      MessageOps.create_message(%{
        conversation_id: conv_id,
        sender: sender_id,
        content: content,
        content_type: content_type,
        recipients: recips,
        file_name: file_name,
        file_size: file_size,
        file_url: file_url,
        status: "pending"
      })

    # ACK to sender
    push(socket, "message:server_ack", %{message_id: message.id, status: "stored"})

    # Create pending receipts
    Enum.each(recips, fn r ->
      MessageReceiptOps.create_receipt(message.id, r)
    end)

    # Broadcast to recipients
    Enum.each(recips, fn r ->
      FundaAfricaWeb.Endpoint.broadcast("chat:user:#{r}", "message:forwarded", %{
        message_id: message.id,
        conversation_id: conv_id,
        from: sender_id,
        content_type: content_type,
        content: content,
        file_name: file_name,
        file_size: file_size,
        file_url: file_url
      })
    end)

    {:noreply, socket}
  end

  def handle_in("message:received", %{"message_id" => mid}, socket) do
    user_id = socket.assigns.user_id

    # Update individual receipt
    MessageReceiptOps.update_status(mid, user_id, "received")

    # Check if all recipients have received the message
    if MessageReceiptOps.all_received?(mid) do
      MessageOps.update_status(mid, "received")
    end

    # Forward acknowledgement back to sender
    sender = MessageOps.get_sender(mid)

    FundaAfricaWeb.Endpoint.broadcast("chat:user:#{sender}", "message:received_ack", %{
      message_id: mid,
      recipient: user_id
    })

    {:noreply, socket}
  end

  def handle_in("message:read", %{"message_id" => mid}, socket) do
    user_id = socket.assigns.user_id

    MessageReceiptOps.update_status(mid, user_id, "read")

    if MessageReceiptOps.all_read?(mid) do
      MessageOps.update_status(mid, "read")
    end

    sender = MessageOps.get_sender(mid)

    FundaAfricaWeb.Endpoint.broadcast("chat:user:#{sender}", "message:read_ack", %{
      message_id: mid,
      recipient: user_id
    })

    {:noreply, socket}
  end

  def handle_in("user:typing", %{"to" => to}, socket) do
    FundaAfricaWeb.Endpoint.broadcast(
      "chat:private:#{socket.assigns.user_id}-#{to}",
      "user:typing",
      %{from: socket.assigns.user_id}
    )

    {:noreply, socket}
  end

  def handle_in("system:announcement", %{"body" => body}, socket) do
    broadcast!(socket, "system:announcement", %{body: body})
    {:noreply, socket}
  end

  def handle_in("message:edit", %{"msg_id" => id, "new_body" => new_body}, socket) do
    broadcast!(socket, "message:edit", %{
      msg_id: id,
      new_body: new_body,
      user: socket.assigns.user_id
    })

    {:noreply, socket}
  end

  def handle_in("message:delete", %{"msg_id" => id}, socket) do
    broadcast!(socket, "message:delete", %{msg_id: id, user: socket.assigns.user_id})
    {:noreply, socket}
  end
end
