defmodule FundaAfrica.Chat.ConversationParticipant do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "conversation_participants" do
    belongs_to :conversation, FundaAfrica.Chat.Conversation
    field :user, :binary_id

    field :role, :string

    timestamps()
  end

  def changeset(participant, attrs) do
    participant
    |> cast(attrs, [:conversation_id, :user, :role])
    |> validate_required([:conversation_id, :user])
  end
end
