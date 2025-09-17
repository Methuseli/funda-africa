defmodule FundaAfricaWeb.ConversationController do
  use FundaAfricaWeb, :controller
  alias FundaAfrica.Chat.ConversationOps

  action_fallback(FundaAfricaWeb.FallbackController)

  # Create private, group, broadcast, or lesson conversation
  def create(
        conn,
        %{"type" => type, "creator" => creator, "participants" => participants} = params
      ) do
    attrs = %{type: type, creator: creator, title: Map.get(params, "title")}
    conv = ConversationOps.create_conversation(attrs, participants)
    conn |> put_status(:created) |> json(%{conversation: conv})
  end

  # Show conversation details
  def show(conn, %{"id" => id}) do
    conv =
      FundaAfrica.Repo.get!(FundaAfrica.Chat.Conversation, id)
      |> FundaAfrica.Repo.preload(:participants)

    json(conn, conv)
  end

  # List all conversations for a user
  def index(conn, params) do
    user_id = params["user_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()

    conversations = ConversationOps.list_by_user(user_id, page, page_size)
    json(conn, conversations)
  end

  def delete(conn, %{"id" => id}) do
    conversation = FundaAfrica.Repo.get!(FundaAfrica.Chat.Conversation, id)
    {:ok, _} = ConversationOps.delete_conversation(conversation)
    send_resp(conn, :no_content, "")
  end
end
