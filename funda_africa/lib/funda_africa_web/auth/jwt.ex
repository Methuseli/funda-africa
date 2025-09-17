defmodule FundaAfricaWeb.Auth.JWT do
  use Joken.Config

  @jwt_secret Application.compile_env(:funda_africa, :jwt_secret)

  @impl true
  def token_config do
    # Skip aud/iss for now; can add later
    default_claims(skip: [:aud, :iss])
  end

  @doc """
  Verifies a JWT token and returns {:ok, claims} or {:error, reason}.
  """
  def verify_token(token) do
    signer = Joken.Signer.create("HS256", @jwt_secret)

    case Joken.verify(token, signer) do
      {:ok, claims} -> {:ok, claims}
      {:error, reason} -> {:error, reason}
    end
  end

  @doc """
  Returns the user_id (subject) if token is valid, otherwise :error.
  """
  def verify_and_get_user(token) do
    case verify_token(token) do
      {:ok, claims} -> {:ok, claims["sub"]}
      _ -> :error
    end
  end
end
