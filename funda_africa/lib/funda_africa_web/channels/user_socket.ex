defmodule FundaAfricaWeb.UserSocket do
  use Phoenix.Socket
  alias FundaAfricaWeb.Auth.JWT

  channel "chat:*", FundaAfricaWeb.ChatChannel
  channel "ai:*", FundaAfricaWeb.AIChannel
  channel "live_lesson:*", FundaAfricaWeb.LiveLessonChannel

  # Connect using cookie or Authorization header
  def connect(params, socket, connect_info) do
    token =
      case get_in(connect_info, [:cookies, "access_token"]) do
        nil -> get_bearer_token(connect_info[:req_headers]) || Map.get(params, "token")
        cookie_token -> cookie_token
      end

    case JWT.verify_and_get_user(token) do
      {:ok, user_id} -> {:ok, assign(socket, :user_id, user_id)}
      :error -> :error
    end
  end

  def connect(_, _, _), do: :error

  def id(socket), do: "users_socket:#{socket.assigns.user_id}"

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
