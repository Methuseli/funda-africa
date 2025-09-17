defmodule FundaAfrica.Chat.ConversationOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.Chat.{Conversation, ConversationParticipant}

  # Create conversation + participants
  def create_conversation(attrs, participants) do
    Repo.transaction(fn ->
      {:ok, conv} =
        %Conversation{}
        |> Conversation.changeset(attrs)
        |> Repo.insert()

      Enum.each(participants, fn user_id ->
        %ConversationParticipant{}
        |> ConversationParticipant.changeset(%{conversation_id: conv.id, user: user_id})
        |> Repo.insert!()
      end)

      conv
    end)
  end

  # Add participant
  def add_participant(conv_id, user_id) do
    %ConversationParticipant{}
    |> ConversationParticipant.changeset(%{conversation_id: conv_id, user: user_id})
    |> Repo.insert()
  end

  # List user conversations (with pagination)
  def list_by_user(user_id, page \\ 1, page_size \\ 20) do
    Conversation
    |> join(:inner, [c], p in assoc(c, :participants))
    |> where([c, p], p.user == ^user_id)
    |> preload([:participants])
    |> order_by(desc: :inserted_at)
    |> Repo.paginate(page: page, page_size: page_size)
  end

  def delete_conversation(%Conversation{} = conversation), do: Repo.delete(conversation)
end
