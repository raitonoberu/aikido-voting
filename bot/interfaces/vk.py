import os
from .abstract import Interface
import vk_api


class Vk(Interface):
    def __init__(self):
        self.api = vk_api.VkApi(token=os.getenv("VK_TOKEN")).get_api()
        self.group_id = os.getenv("VK_GROUPID")

    def send(self, text: str):
        self.api.wall.post(message=text, from_group=True, owner_id=self.group_id)
