from pytube import YouTube 
import ssl
import librosa
import soundfile as sf
import torch
import librosa
from transformers import Wav2Vec2ForCTC, Wav2Vec2Processor

ssl._create_default_https_context = ssl._create_unverified_context

restriction = False

# link = "https://www.youtube.com/watch?v=9l5j6v4-x5o"
# yt = YouTube(link) 

# try: 
#     yt.streams.filter(res="360p").first().download(filename="video.mp4")
# except: 
#     restriction = True
#     print("age restricted video") 


def extract_audio_from_video(video_path,audio_path):
   audio , sr = librosa.load(video_path)
   sf.write(audio_path,audio,sr)

   speech_to_text(audio_path)

def speech_to_text(audio_path):
    MODEL_ID = "jonatasgrosman/wav2vec2-large-xlsr-53-russian"
    AUDIO_FILE_PATH = audio_path

    # Load the pretrained processor and model
    processor = Wav2Vec2Processor.from_pretrained(MODEL_ID)
    model = Wav2Vec2ForCTC.from_pretrained(MODEL_ID)

    # Load the audio file
    speech_array, sampling_rate = librosa.load(AUDIO_FILE_PATH, sr=16_000)

    # Prepare the input
    inputs = processor(speech_array, sampling_rate=16_000, return_tensors="pt", padding=True)

    # Perform the inference
    with torch.no_grad():
        logits = model(inputs.input_values, attention_mask=inputs.attention_mask).logits

    # Decode the predicted IDs to text
    predicted_ids = torch.argmax(logits, dim=-1)
    predicted_sentence = processor.batch_decode(predicted_ids)[0]

    # Print the prediction
    print("-" * 100)
    print("Text:", predicted_sentence)
    print("-" * 100)

    words = predicted_sentence.split();
    # themeKeyWords = [
    # "Казино": ["Лас Вегас", "игровые автоматы", "ставки", "джекпот", "выигрыш", "рулетка", "блэкджек", "азартные игры", "фишки", "казино-игры", "спустила деньги", "игровой стол", "игровой зал"],
    # "Наука": ["порошок", "очистка воды", "вода", "чистая вода", "фильтрация", "очищение", "водный фильтр", "демонстрация", "эксперимент", "наука", "технология"],
    # "Насилие": ["драка", "удары", "агрессия", "конфликт", "нападение", "кровь", "насилие", "жестокость", "потасовка", "ссора"],
    # "Обучение": ["учить", "математика", "считать", "посчитаем", "задание", "найди", "сколько", "плюс", "минус", "равно"],
    # "Кино": ["обзор", "фильмы", "топ"],
    # "Настольные игры": ["подземелья", "приключения", "ролевые игры", "игровая сессия", "квест", "кампания", "гейммастер", "игроки", "правила игры", "настольная игра"],
    # "Видеоигры": ["геймплей", "уровни", "боссы", "квесты", "миссии", "очки опыта", "инвентарь", "сюжетная линия", "графика", "мультиплеер", "обновления", "игровой процесс", "контроллеры", "персонажи", "режимы игры"],
    # "Детские мультики": ["анимация", "мультфильм", "детский", "персонажи", "сказка", "приключения", "обучение", "доброта", "дружба", "животные", "волшебство", "цветной", "песни", "веселый", "семейный"]
    # ]
    themeKeyWords = [["Лас Вегас", "ставки", "джекпот", "выигрыш", "казино-игры"]]
    themes = [0, 0, 0, 0, 0, 0, 0, 0]

    def dist(a, b):
        def rec(i, j):
            if i == 0 or j == 0: 
                return max(i, j)
            elif a[i-1] == b[j-1]:
                return rec(i-1, j-1)
            else:
                return 1 + min(rec(i, j-1),
                               rec(i-1, j),
                               rec(i-1, j-1))

        return rec(len(a), len(b))

    for word in words:
        for idx, theme in enumerate(themeKeyWords):
            for keyWord in theme:
                lev = dist(keyWord, word)
                bigger = max([len(keyWord), len(word)])
                pct = ((bigger - lev) / bigger) * 100
                if(pct >= 50):
                    themes[idx] = themes[idx]+1
                    print('Строка #1: {str1}\nСтрока #2: {str2}\n=============\nСхожесть: {pct}%'.format(str1 = keyWord, str2 = word, pct = pct))
                    print('\n')   

    for theme in themes:
        print(theme)

# if(restriction == False):
#     extract_audio_from_video("video.mp4","audio.wav")
# else:
#     print("Restricted Content")       

speech_to_text("audio.ogg")