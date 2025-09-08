defmodule FundaAfrica.AI.AIInteraction do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "ai_interactions" do
    belongs_to :user, FundaAfrica.Accounts.User

    field :prompt, :string
    field :response, :string
    field :model_used, :string
    field :metadata, :map

    timestamps()
  end

  def changeset(interaction, attrs) do
    interaction
    |> cast(attrs, [:user_id, :prompt, :response, :model_used, :metadata])
    |> validate_required([:user_id, :prompt, :response])
  end
end
