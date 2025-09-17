defmodule FundaAfricaWeb.Auth.JWT do
  use Joken.Config

  @impl true
  def token_config do
    default_claims(skip: [:aud, :iss])
  end
end
