defmodule FundaAfrica.LiveLessons.Invite do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "lesson_invites" do
    belongs_to :lesson, FundaAfrica.LiveLessons.LiveLesson

    field :is_active, :boolean, default: true
    field :metadata, :map, default: %{}

    timestamps()
  end

  def changeset(invite, attrs) do
    invite
    |> cast(attrs, [:lesson_id, :status, :metadata])
    |> validate_required([:lesson_id])
  end
end
