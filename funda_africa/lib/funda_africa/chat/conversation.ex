defmodule FundaAfrica.Chat.Conversation do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "conversations" do
    field :title, :string
    field :type, :string

    field :creator, :binary_id

    has_many :participants, FundaAfrica.Chat.ConversationParticipant
    has_many :messages, FundaAfrica.Chat.Message

    timestamps()
  end

  def changeset(conversation, attrs) do
    conversation
    |> cast(attrs, [:title, :type, :creator])
    |> validate_required([:type, :creator])
  end
end
