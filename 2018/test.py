from collections import Counter
from wordcloud import WordCloud


def wc(name):
    txt = ""
    with open("/mnt/d/input.txt", "r") as f:
        for line in f:
            try:
                [_, user, msg] = line.strip().split(":", 2)
                txt += (
                    (msg.replace("<Media omitted>", "") + " ") if name in user else ""
                )
            except:
                pass
    sw = [
        "estaba",
        "https",
        "pero",
        "como",
        "xddd",
        "xdxd",
        "xdxdx",
        "xdxdxd",
        "para",
        "xddd",
        "xdddd",
        "xddddd",
        "XDDDDD",
    ]
    WordCloud(width=1920, height=1080, stopwords=sw, min_word_length=10).generate(
        txt
    ).to_file(f"/mnt/d/{name}.png")


for name in ["Daniel", "Doctora", "Alejandro"]:
    wc(name)
