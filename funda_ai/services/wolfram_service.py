# services/wolfram_service.py
import httpx
from typing import Any, Dict
import xml.etree.ElementTree as ET


class WolframError(Exception):
    """Custom exception for Wolfram API errors."""
    pass


class WolframService:
    """
    Service class to interact with the Wolfram Alpha Full Results API,
    handling the XML response and converting it to a dictionary.
    """

    def __init__(
        self,
        appid: str,
        client: httpx.AsyncClient,
        base_url: str = "https://api.wolframalpha.com/v2",
    ):
        """
        Initializes the Wolfram service with an AppID and an async HTTP client.
        
        :param appid: The AppID for the Wolfram Alpha service.
        :param client: An httpx.AsyncClient instance.
        :param base_url: The base URL for the Wolfram Alpha API.
        """
        self.appid = appid
        self.client = client
        self.base_url = base_url

    async def query(self, input_text: str) -> Dict[str, Any]:
        """
        Queries Wolfram Alpha and returns the XML response as a dictionary.
        
        :param input_text: The query string for Wolfram Alpha.
        :return: A dictionary representation of the Wolfram Alpha XML response.
        :raises WolframError: If the API request or XML parsing fails.
        """
        params = {
            "appid": self.appid,
            "input": input_text,
            # Wolfram Alpha returns XML by default
        }
        try:
            resp = await self.client.get(f"{self.base_url}/query", params=params)
            resp.raise_for_status()

            # The Wolfram Alpha API returns XML, so we parse the response text
            xml_text = resp.text
            root = ET.fromstring(xml_text)
            
            # Convert the parsed XML tree to a dictionary
            return self._xml_to_dict(root)
        
        except httpx.HTTPError as exc:
            raise WolframError(f"Wolfram request failed: {exc}") from exc
        except ET.ParseError as exc:
            raise WolframError(f"Failed to parse XML response from Wolfram: {exc}") from exc
        except Exception as exc:
            raise WolframError(f"An unexpected error occurred: {exc}") from exc

    def _xml_to_dict(self, element: ET.Element) -> Dict[str, Any]:
        """
        Recursively converts an XML element tree to a Python dictionary.
        
        :param element: The XML element to convert.
        :return: A dictionary representing the XML element.
        """
        if len(element) == 0 and not element.text:
            return {}

        d = {}
        for child in element:
            child_dict = self._xml_to_dict(child)
            if child.tag in d:
                if not isinstance(d[child.tag], list):
                    d[child.tag] = [d[child.tag]]
                d[child.tag].append(child_dict)
            else:
                d[child.tag] = child_dict

        if element.text and element.text.strip():
            d["#text"] = element.text.strip()
            
        if element.attrib:
            d["@attributes"] = element.attrib

        return d
