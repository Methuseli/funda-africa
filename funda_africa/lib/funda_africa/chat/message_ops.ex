defmodule FundaAfrica.Chat.MessageOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.Chat.Message

  def list_messages, do: Repo.all(Message)
  def get_message!(id), do: Repo.get!(Message, id)

  def create_message(attrs \\ %{}) do
    %Message{}
    |> Message.changeset(attrs)
    |> Repo.insert()
  end

  def update_message(%Message{} = message, attrs) do
    message
    |> Message.changeset(attrs)
    |> Repo.update()
  end

  def delete_message(%Message{} = message), do: Repo.delete(message)

  def list_by_conversation(conversation_id) do
    Message
    |> where(conversation_id: ^conversation_id)
    |> Repo.all()
  end
end
