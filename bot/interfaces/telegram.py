import os
from .abstract import Interface
import telebot


class Telegram(Interface):
    def __init__(self):
        self.api = telebot.TeleBot(os.getenv("TG_TOKEN"))
        self.chat_id = os.getenv("TG_CHATID")

    def send(self, text: str):
        self.api.send_message(self.chat_id, text)
