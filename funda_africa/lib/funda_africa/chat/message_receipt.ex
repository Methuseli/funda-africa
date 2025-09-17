defmodule FundaAfrica.Chat.MessageReceipt do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "message_receipts" do
    belongs_to :message, FundaAfrica.Chat.Message, type: :binary_id
    field :recipient_id, :binary_id
    field :status, :string, default: "pending" # pending, received, read

    timestamps()
  end

  def changeset(receipt, attrs) do
    receipt
    |> cast(attrs, [:message_id, :recipient_id, :status])
    |> validate_required([:message_id, :recipient_id])
  end
end
