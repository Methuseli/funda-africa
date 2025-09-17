defmodule FundaAfricaWeb.UserSocket do
  use Phoenix.Socket
  alias FundaAfricaWeb.Auth.JWT

  ## Channels
  channel "chat:*", FundaAfricaWeb.ChatChannel
  channel "ai:*", FundaAfricaWeb.AIChannel

  # Connect: handles cookie or Authorization header
  def connect(params, socket, connect_info) do
    token =
      # Check cookie
      case get_in(connect_info, [:cookies, "access_token"]) do
        nil ->
          # Check Authorization header (mobile)
          get_bearer_token(connect_info[:req_headers]) || Map.get(params, "token")
        cookie_token ->
          cookie_token
      end

    case verify_jwt(token) do
      {:ok, user_id} -> {:ok, assign(socket, :user_id, user_id)}
      _ -> :error
    end
  end

  def connect(_, _, _), do: :error

  # Socket id for presence
  def id(socket), do: "users_socket:#{socket.assigns.user_id}"

  # JWT verification
  defp verify_jwt(nil), do: :error
  defp verify_jwt(token) do
    case JWT.verify_and_validate(token) do
      {:ok, claims} -> {:ok, claims["sub"]}
      _ -> :error
    end
  end

  # Extract bearer token from headers
  defp get_bearer_token(headers) when is_list(headers) do
    headers
    |> Enum.find_value(nil, fn
      {"authorization", "Bearer " <> token} -> token
      _ -> nil
    end)
  end

  defp get_bearer_token(_), do: nil
end
