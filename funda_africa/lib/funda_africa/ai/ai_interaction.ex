defmodule FundaAfrica.AI.AIInteraction do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "ai_interactions" do
    belongs_to :conversation, FundaAfrica.Chat.Conversation
    field :user, :binary_id

    field :prompt, :string
    field :response, :string
    field :model_used, :string
    field :metadata, :map
    field :topic, :string

    timestamps()
  end

  def changeset(interaction, attrs) do
    interaction
    |> cast(attrs, [:user, :prompt, :response, :model_used, :metadata])
    |> validate_required([:user, :prompt, :response])
  end
end
