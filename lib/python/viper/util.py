def verify_url(url: str) -> bool:
    if not url.startswith("http://") and not url.startswith("https://"):
        return False
    return True
