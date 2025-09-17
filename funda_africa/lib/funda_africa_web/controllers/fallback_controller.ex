defmodule FundaAfricaWeb.FallbackController do
  use FundaAfricaWeb, :controller

  def call(conn, {:error, %Ecto.Changeset{} = changeset}) do
    conn
    |> put_status(:unprocessable_entity)
    |> json(%{errors: Ecto.Changeset.traverse_errors(changeset, &translate_error/1)})
  end

  defp translate_error({msg, opts}), do: Phoenix.Naming.humanize(msg)
end
