defmodule FundaAfrica.Lessons.Invite do
  use Ecto.Schema
  import Ecto.Changeset

  @primary_key {:id, :binary_id, autogenerate: true}
  @foreign_key_type :binary_id
  schema "lesson_invites" do
    belongs_to :lesson, FundaAfrica.Lessons.Lesson

    field :is_active, :boolean, default: true
    field :status, :string, default: "pending" # pending | accepted | declined
    field :metadata, :map, default: %{}

    timestamps()
  end

  def changeset(invite, attrs) do
    invite
    |> cast(attrs, [:lesson_id, :inviter_id, :invitee_id, :status, :metadata])
    |> validate_required([:lesson_id, :inviter_id, :invitee_id])
    |> validate_inclusion(:status, ["pending", "accepted", "declined"])
  end
end


  def list_for_user(conn, params) do
    user_id = params["user_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()
    invites = LessonInviteOps.list_for_user(user_id, page, page_size)
    json(conn, invites)
  end

  def accept(conn, %{"id" => invite_id}) do
    case LessonInviteOps.accept_invite(invite_id) do
      {:ok, invite} ->
        json(conn, %{invite: invite})

      {:error, reason} ->
        conn |> put_status(:bad_request) |> json(%{error: reason})
    end
  end

  def decline(conn, %{"id" => invite_id}) do
    case LessonInviteOps.decline_invite(invite_id) do
      {:ok, invite} ->
        json(conn, %{invite: invite})

      {:error, reason} ->
        conn |> put_status(:unprocessable_entity

  def list_for_user(conn, params) do
    user_id = params["user_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()
    invites = LessonInviteOps.list_for_user(user_id, page, page_size)
    json(conn, invites)
  end

  def accept(conn, %{"id" => invite_id}) do
    case LessonInviteOps.accept_invite(invite_id)

  def list_for_user(conn, params) do
    user_id = params["user_id"]
    page = Map.get(params, "page", "1") |> String.to_integer()
    page_size = Map.get(params, "page_size", "20") |> String.to_integer()
    invites = LessonInviteOps.list_for_user(user_id, page, page_size)
    json(conn, invites)
  end

  def accept(conn, %{"id" => invite_id}) do
    case LessonInviteOps.accept_invite(invite_id) do
      {:ok, invite} ->
        json(conn, %{invite: invite})

      {:error, reason} ->
     do
      {:ok, invite} ->
        json(conn, %{invite: invite})

      {:error, reason} ->
        conn |> put_status(:bad_request) |> json(%{error: reason})
    end
  end
end
