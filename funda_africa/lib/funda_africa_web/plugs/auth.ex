defmodule FundaAfricaWeb.Plugs.Auth do
  import Plug.Conn
  alias FundaAfricaWeb.Auth.JWT

  def init(opts), do: opts

  def call(conn, _opts) do
    case get_token(conn) do
      nil ->
        conn
        |> send_resp(401, "Unauthorized")
        |> halt()

      token ->
        case JWT.verify_and_get_user(token) do
          {:ok, user_id} ->
            assign(conn, :current_user, user_id)

          :error ->
            conn
            |> send_resp(401, "Unauthorized")
            |> halt()
        end
    end
  end

  defp get_token(conn) do
    conn = fetch_cookies(conn)

    cond do
      conn.cookies["user_token"] ->
        conn.cookies["user_token"]

      auth_header = get_req_header(conn, "authorization") |> List.first() ->
        case auth_header do
          "Bearer " <> token -> token
          _ -> nil
        end

      true ->
        nil
    end
  end
end
