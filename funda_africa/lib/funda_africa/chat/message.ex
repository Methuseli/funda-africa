defmodule FundaAfrica.Chat.Message do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "messages" do
    belongs_to :conversation, FundaAfrica.Chat.Conversation
    belongs_to :sender, FundaAfrica.Accounts.User

    field :content, :string
    field :content_type, :string, default: "text"

    timestamps()
  end

  def changeset(message, attrs) do
    message
    |> cast(attrs, [:conversation_id, :sender_id, :content, :content_type])
    |> validate_required([:conversation_id, :sender_id, :content])
  end
end
