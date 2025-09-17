defmodule FundaAfricaWeb.ConversationParticipantController do
  use FundaAfricaWeb, :controller
  alias FundaAfrica.Chat.ConversationOps

  action_fallback FundaAfricaWeb.FallbackController

  def create(conn, %{"conversation_id" => conv_id, "user_id" => user_id}) do
    case ConversationOps.add_participant(conv_id, user_id) do
      {:ok, participant} -> json(conn, %{participant: participant})
      {:error, changeset} -> conn |> put_status(:unprocessable_entity) |> json(%{errors: changeset})
    end
  end
end
