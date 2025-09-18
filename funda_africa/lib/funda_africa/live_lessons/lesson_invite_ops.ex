defmodule FundaAfrica.Lessons.InviteOps do
  @moduledoc """
  Operations for managing lesson invites.
  """

  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.Lessons.Invite

  # Create a new invite
  def create_invite(attrs) do
    %Invite{}
    |> Invite.changeset(attrs)
    |> Repo.insert()
  end

  # Update an invite (e.g., accept/decline)
  def update_invite(%Invite{} = invite, attrs) do
    invite
    |> Invite.changeset(attrs)
    |> Repo.update()
  end

  # Delete an invite
  def delete_invite(%Invite{} = invite) do
    Repo.delete(invite)
  end

  # Get a single invite by ID
  def get_invite!(id), do: Repo.get!(Invite, id)

  # List all invites for a given lesson
  def list_by_lesson(lesson_id) do
    Invite
    |> where([i], i.lesson_id == ^lesson_id)
    |> Repo.all()
  end

  # List all invites for a given user (either inviter or invitee)
  def list_for_user(user_id, page \\ 1, page_size \\ 20) do
    Invite
    |> where([i], i.invitee_id == ^user_id, i.is_active == true)
    |> Repo.paginate(page: page, page_size: page_size)
  end

  # Accept an invite
  def accept_invite(%Invite{} = invite) do
    update_invite(invite, %{status: "accepted"})
  end

  # Decline an invite
  def decline_invite(%Invite{} = invite) do
    update_invite(invite, %{status: "declined"})
  end
end
