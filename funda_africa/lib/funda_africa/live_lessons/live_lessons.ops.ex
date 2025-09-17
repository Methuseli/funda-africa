defmodule FundaAfrica.LiveLessons.LiveLessonOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.LiveLessons.LiveLesson

  def list_live_lessons, do: Repo.all(LiveLesson)
  def get_live_lesson!(id), do: Repo.get!(LiveLesson, id)

  def create_live_lesson(attrs \\ %{}) do
    %LiveLesson{}
    |> LiveLesson.changeset(attrs)
    |> Repo.insert()
  end

  def update_live_lesson(%LiveLesson{} = lesson, attrs) do
    lesson
    |> LiveLesson.changeset(attrs)
    |> Repo.update()
  end

  def delete_live_lesson(%LiveLesson{} = lesson), do: Repo.delete(lesson)

  def list_by_host(host_id) do
    LiveLesson
    |> where(host: ^host_id)
    |> Repo.all()
  end

  @doc """
  Returns true if the lesson is joinable (scheduled or active).
  """
  def lesson_active?(lesson_id) do
    case Repo.get(LiveLesson, lesson_id) do
      %LiveLesson{status: status} when status in ["scheduled", "active"] -> true
      _ -> false
    end
  end
end
