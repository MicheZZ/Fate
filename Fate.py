from random  import randint
print("Задай свой вопрос. Примечание: вопрос должен заканчиваться '?'")
print("Можешь указать к какому ответу больше склоняетесь. Будешь?")

key2 = None
key3 = None
while key2 != True:
    answer = input()
    if answer.lower() != "да" and answer.lower() != "нет":
        print("Напиши да или нет")
    else:
        if answer.lower() == "да":
            print("Хорошо. Выбери к какому ответу склоняешься больше всего")
            print("Нужно выбрать целое число от 1 до 99, где 0 - нет, а 100 - да")
            while key3 != True:
                try:
                    answer_pr = int(input())
                    if answer_pr >= 1 and answer_pr <= 99:
                        key3 = True
                        key2 = True
                    else:
                        print("Введи целое число от 1 до 100")
                except:
                    print("Введи целое число от 1 до 100")
        else:
            key2 = True

print("Задавай вопрос")
d = randint(0, 3)
l = randint(0, 4)
positive = ["Бесспорно", "Предрешено", "Никаких сомнений", "Определённо да", "Можешь быть уверен в этом"]
indecpositive = ["Мне кажется - да","Вероятнее всего","Хорошие перспективы","Знаки говорят - да","Да"]
neutral = ["Пока неясно, попробуй снова","Спроси позже","Лучше не рассказывать","Сейчас нельзя предсказать","Сконцентрируйся и спроси опять"]
negative = ["Даже не думай","Мой ответ - нет","По моим данным - нет","Перспективы не очень хорошие","Весьма сомнительно"]
s = [positive, indecpositive, neutral, negative]




#ввод
key = None
while key != 100:
    g = input()
    try:
        g = int(g)
        print("Задайте свой вопрос корректно. Примечание: вопрос должен заканчиваться '?'")
    except ValueError:
        key = True
        # вывод
        if key == True and g[-1] == "?":
            print(s[d][l])
            break
        else:
            print("Задайте свой вопрос корректно. Примечание: вопрос должен заканчиваться '?'")
input()