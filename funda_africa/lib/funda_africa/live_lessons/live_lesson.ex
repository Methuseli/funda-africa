defmodule FundaAfrica.LiveLessons.LiveLesson do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "live_lessons" do
    field :host, :binary_id

    field :title, :string
    field :description, :string
    field :start_time, :utc_datetime
    field :end_time, :utc_datetime
    field :status, :string, default: "scheduled"

    timestamps()
  end

  def changeset(lesson, attrs) do
    lesson
    |> cast(attrs, [:host, :title, :description, :start_time, :end_time, :status])
    |> validate_required([:host, :title, :start_time])
  end
end
