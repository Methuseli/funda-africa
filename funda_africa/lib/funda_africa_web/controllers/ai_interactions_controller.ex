defmodule FundaAfricaWeb.AIInteractionController do
  use FundaAfricaWeb, :controller
  alias FundaAfrica.Chat.ConversationOps
  alias FundaAfrica.AI.AIInteractionOps

  def index(conn, params) do
    user_id = params["user_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer
    conversations = ConversationOps.list_user_ai_conversations(user_id, page, page_size)

    json(conn, conversations)
  end

  def show(conn, params) do
    conversation_id = params["conversation_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()
    interactions = AIInteractionOps.list_by_conversation(conversation_id, page, page_size)

    json(conn, interactions)
  end
end
