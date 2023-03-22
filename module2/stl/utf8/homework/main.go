package main

import (
	"fmt"
	"github.com/alexsergivan/transliterator"
	"strings"
	"unicode/utf8"
)

var data = ` Сообщения с телеграм канала Go-go!
я, извините, решил, что вы хоть что-то про регулярки знаете 🙂
С Почтением, с уважением 🙂.

Сообщения с телеграм канала Go-get job:
#Вакансия #Vacancy #Удаленно #Remote #Golang #Backend  
👨‍💻  Должность: Golang Tech Lead
🏢  Компания: PointPay
💵  Вилка: от 9000 до 15000 $ до вычета налогов
🌏  Локация: Великобритания
💼  Формат работы: удаленный
💼  Занятость: полная занятость 

Наша компания развивает 15 продуктов в сфере блокчейн и криптовалюты, функционирующих в рамках единой экосистемы. Мы - одни из лучших в своей нише, и за без малого три года мы создали криптовалютную платформу, которая предоставляет комплексные решения и позволяет пользователям проводить практически любые финансовые операции.
Наш проект победил в номинации "лучший блокчейн-стартап 2019 года" на одной из самых крупных конференций по блокчейн, криптовалютам и майнингу – Blockchain Life в Москве.

У нас очень амбициозные планы по развитию наших продуктов, и на данный момент мы находимся в поиске Golang Technical Lead, чтобы вместе создать революционный и самый технологичный продукт в мире.


  
Мы ожидаем от тебя:  
  
✅ Опыт управления финансовыми IT продуктами;
✅ Опыт постановки задач технической команде и приема результатов;
✅ Понимание и использование современных подходов к разработке (Agile);
✅ Опыт разработки не менее 5 лет;
✅ Отличное владение Go, будет плюсом опыт работы на PHP;
✅ Опыт работы с REST / GRPС (protobuf);
✅ Опыт работы с Laravel;
✅ Отличные навыки SQL (мы используем PostgreSQL);
✅ Опыт работы с Redis или аналогичной системой кеширования;
✅ Опыт работы с Kubernetes, с брокерами сообщений (RabbitMQ, Kafka и др.);
✅ Базовые знания Amazon Web Services.

Будем плюсом:

✅ Опыт разработки криптовалютных кошельков / криптовалютных бирж / криптовалютных торговых роботов / криптовалютных платежных систем или опыт работы в FinTech или банковских продуктах.

Чем предстоит заниматься:  
  
📌 Управлением и взаимодействием с командой;
📌 Постановкой задач команде разработчиков;
📌 Разработкой решений для запуска и сопровождения разрабатываемого компанией ПО;
📌 Выстраиванием эффективного процесса разработки;
📌 Достигать результатов, поддерживать и развивать текущие проекты, развивать новые проекты и продукты, искать и предлагать нестандартные решения задач.

Мы гарантируем будущему коллеге:  
  
🔥 Высокую заработную плату до 15 000$ (в валюте);
🔥 Индексацию заработной платы;
🔥 Полностью удаленный формат работы и максимально гибкий график - работайте из любой точки мира в удобное вам время;
🔥 Оформление по бессрочному международному трудовому договору с UK;
🔥 Отсутствие бюрократии и кучи бессмысленных звонков;
🔥 Корпоративную технику за счет компании;
🔥 Сотрудничество с командой высококвалифицированных специалистов, настоящих профессионалов своего дела;
🔥 Работу над масштабным, интересным и перспективным проектом.


Если в этом описании вы узнали  себя - пишите. Обсудим проект, задачи и все детали :)  
  
Павел,  
✍️TG: @pavel_hr  
📫 papolushkin.wanted@gmail.com`

func main() {
	fmt.Println("Кол-во символов в тексте:", utf8.RuneCountInString(data))
	data = strings.ReplaceAll(data, "🙂", "=)")
	CyrillicByteCount := utf8.RuneCountInString(data)
	fmt.Println("Кол-во символов после замены смайлов:", CyrillicByteCount)
	translit := transliterator.NewTransliterator(nil)
	trans := translit.Transliterate(data, "en")
	LatinByteCount := utf8.RuneCountInString(trans)
	fmt.Println("Кол-во символов после транслита:", LatinByteCount)
	res := float64(LatinByteCount) / float64(CyrillicByteCount)
	fmt.Println("Компрессия:", res)
}
