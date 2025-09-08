defmodule FundaAfrica.Chat.Conversation do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "conversations" do
    field :title, :string
    field :type, :string

    belongs_to :creator, FundaAfrica.Accounts.User, foreign_key: :created_by

    has_many :participants, FundaAfrica.Chat.ConversationParticipant
    has_many :messages, FundaAfrica.Chat.Message

    timestamps()
  end

  def changeset(conversation, attrs) do
    conversation
    |> cast(attrs, [:title, :type, :created_by])
    |> validate_required([:type, :created_by])
  end
end
