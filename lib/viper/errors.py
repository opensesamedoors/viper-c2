class NotAuth(Exception):
    def __init__(self):
        super().__init__("Operator is not authenticated")

class CookieNotFound(Exception):
    def __init__(self):
        super().__init__("Server did not set a cookie")

class ServerError(Exception):
    pass
