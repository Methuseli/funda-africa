defmodule FundaAfrica.LiveLessons.LessonParticipant do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "lesson_participants" do
    belongs_to :lesson, FundaAfrica.LiveLessons.LiveLesson
    field :user, :binary_id

    field :role, :string, default: "student"

    timestamps()
  end

  def changeset(participant, attrs) do
    participant
    |> cast(attrs, [:lesson_id, :user, :role])
    |> validate_required([:lesson_id, :user])
  end
end
