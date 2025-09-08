defmodule FundaAfrica.Chat.ConversationParticipantOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.Chat.ConversationParticipant

  @doc "List all conversation participants"
  def list_conversation_participants do
    Repo.all(ConversationParticipant)
  end

  @doc "Get a conversation participant by ID (raises if not found)"
  def get_conversation_participant!(id), do: Repo.get!(ConversationParticipant, id)

  @doc "Create a conversation participant"
  def create_conversation_participant(attrs \\ %{}) do
    %ConversationParticipant{}
    |> ConversationParticipant.changeset(attrs)
    |> Repo.insert()
  end

  @doc "Update a conversation participant"
  def update_conversation_participant(%ConversationParticipant{} = participant, attrs) do
    participant
    |> ConversationParticipant.changeset(attrs)
    |> Repo.update()
  end

  @doc "Delete a conversation participant"
  def delete_conversation_participant(%ConversationParticipant{} = participant) do
    Repo.delete(participant)
  end

  @doc "Get all participants in a conversation"
  def list_by_conversation(conversation_id) do
    ConversationParticipant
    |> where(conversation_id: ^conversation_id)
    |> Repo.all()
  end
end
