defmodule FundaAfrica.Chat.MessageReceiptOps do
  import Ecto.Query, warn: false
  alias FundaAfrica.Repo
  alias FundaAfrica.Chat.MessageReceipt

  @doc """
  Creates a receipt for a given message and recipient.
  """
  def create_receipt(message_id, recipient_id) do
    %MessageReceipt{}
    |> MessageReceipt.changeset(%{message_id: message_id, recipient_id: recipient_id})
    |> Repo.insert()
  end

  @doc """
  Updates the status of a receipt (e.g., "received", "read").
  """
  def update_status(message_id, recipient_id, status) do
    case Repo.get_by(MessageReceipt, message_id: message_id, recipient_id: recipient_id) do
      nil ->
        {:error, :not_found}

      receipt ->
        receipt
        |> Ecto.Changeset.change(status: status)
        |> Repo.update()
    end
  end

  @doc """
  Returns all receipts for a message.
  """
  def list_for_message(message_id) do
    from(r in MessageReceipt, where: r.message_id == ^message_id)
    |> Repo.all()
  end

  @doc """
  Checks if all recipients have acknowledged (e.g. all `status = "received"`).
  """
  def all_received?(message_id) do
    from(r in MessageReceipt,
      where: r.message_id == ^message_id and r.status != "received",
      select: count(r.id)
    )
    |> Repo.one()
    |> Kernel.==(0)
  end

  def all_read?(message_id) do
    from(r in MessageReceipt,
      where: r.message_id == ^message_id and r.status != "read",
      select: count(r.id)
    )
    |> Repo.one()
    |> Kernel.==(0)
  end
end
