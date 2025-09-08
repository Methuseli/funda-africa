defmodule FundaAfrica.Chat.ConversationOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.Chat.Conversation

  def list_conversations, do: Repo.all(Conversation)
  def get_conversation!(id), do: Repo.get!(Conversation, id)

  def create_conversation(attrs \\ %{}) do
    %Conversation{}
    |> Conversation.changeset(attrs)
    |> Repo.insert()
  end

  def update_conversation(%Conversation{} = conversation, attrs) do
    conversation
    |> Conversation.changeset(attrs)
    |> Repo.update()
  end

  def delete_conversation(%Conversation{} = conversation), do: Repo.delete(conversation)

  def list_by_creator(user_id) do
    Conversation
    |> where(created_by: ^user_id)
    |> Repo.all()
  end
end
