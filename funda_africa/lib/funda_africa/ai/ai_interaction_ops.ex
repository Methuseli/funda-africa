defmodule FundaAfrica.AI.AIInteractionOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.AI.AIInteraction

  def list_ai_interactions, do: Repo.all(AIInteraction)
  def get_ai_interaction!(id), do: Repo.get!(AIInteraction, id)

  def create_ai_interaction(attrs \\ %{}) do
    %AIInteraction{}
    |> AIInteraction.changeset(attrs)
    |> Repo.insert()
  end

  def update_ai_interaction(%AIInteraction{} = interaction, attrs) do
    interaction
    |> AIInteraction.changeset(attrs)
    |> Repo.update()
  end

  def delete_ai_interaction(%AIInteraction{} = interaction), do: Repo.delete(interaction)

  def list_by_user(user_id) do
    AIInteraction
    |> where(user_id: ^user_id)
    |> Repo.all()
  end

  def list_by_conversation(conversation_id, page \\ 1, page_size \\ 20) do
    AIInteraction
    |> where(conversation_id: ^conversation_id)
    |> Repo.paginate(page: page, page_size: page_size)
  end
end
