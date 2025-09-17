defmodule FundaAfricaWeb.MessageController do
  use FundaAfricaWeb, :controller
  alias FundaAfrica.Chat.MessageOps

  action_fallback FundaAfricaWeb.FallbackController

  # Get paginated messages for a conversation
  def index(conn, %{"id" => conv_id, "page" => page, "page_size" => page_size}) do
    messages = MessageOps.list_by_conversation(conv_id, String.to_integer(page), String.to_integer(page_size))
    json(conn, messages)
  end
end
