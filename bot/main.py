from datetime import datetime
import time
from service import Service
from interfaces import Interface, Vk, Telegram

template = """{group}, новый {type} опрос!

Название: {name}
Описание: {description}

Доступен до {expires_at}"""


def format_pool(pool) -> str:
    return template.format(
        group=pool["group"]["name"],
        type="анонимный" if pool["is_anonymous"] else "публичный",
        name=pool["name"],
        description=pool["description"],
        expires_at=datetime.fromisoformat(pool["expires_at"]).strftime(
            "%d.%m.%Y, %H:%M"
        ),
    )


service = Service()
interfaces: list[Interface] = [Vk(), Telegram()]


def main():
    last_id = None
    while True:
        try:
            pools = service.get_pools()
        except Exception as e:
            print("Service:", e)
            time.sleep(5)
            continue

        if len(pools) == 0 or last_id == pools[0]["id"]:
            time.sleep(60)
            continue
        if last_id == None:
            last_id = pools[0]["id"]
            continue

        print("New pool:", pools[0]["id"])

        text = format_pool(pools[0])
        for i in interfaces:
            try:
                i.send(text)
            except Exception as e:
                print(type(i).__name__ + ":", e)
        last_id = pools[0]["id"]


if __name__ == "__main__":
    main()
