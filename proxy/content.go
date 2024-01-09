package main

const (
	swaggerTemplate = `<!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <meta http-equiv="X-UA-Compatible" content="ie=edge">
        <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-standalone-preset.js"></script>
        <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-standalone-preset.js"></script> -->
        <script src="//unpkg.com/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
        <!-- <script src="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui-bundle.js"></script> -->
        <link rel="stylesheet" href="//unpkg.com/swagger-ui-dist@3/swagger-ui.css" />
        <!-- <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/swagger-ui/3.22.1/swagger-ui.css" /> -->
        <style>
            body {
                margin: 0;
            }
        </style>
        <title>Swagger</title>
    </head>
    <body>
        <div id="swagger-ui"></div>
        <script>
            window.onload = function() {
              SwaggerUIBundle({
                url: "/docs/swagger.json?{{.Time}}",
                dom_id: '#swagger-ui',
                presets: [
                  SwaggerUIBundle.presets.apis,
                  SwaggerUIStandalonePreset
                ],
                layout: "StandaloneLayout"
              })
            }
        </script>
    </body>
    </html>
    `
)

const content1 = `---
menu:
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: ` + "`" + `/app/static/tasks/_index.md` + "`" + `

Должен меняться счетчик и время:

Текущее время:` // 2021-10-13 15:00:00

const content2 = `
Счетчик:` // 0

const content3 = `
## Критерии приемки:
- [ ] Воркер должен обновлять данные каждые 5 секунд
- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд
- [ ] Время должно обновляться каждые 5 секунд`

const content4 = `---
menu:
    after:
        name: binary_tree
        weight: 2
title: Построение сбалансированного бинарного дерева
---

# Задача построить сбалансированное бинарное дерево
Используя AVL дерево, постройте сбалансированное бинарное дерево, на текущей странице.

Нужно написать воркер, который стартует дерево с 5 элементов, и каждые 5 секунд добавляет новый элемент в дерево.

Каждые 5 секунд на странице появляется актуальная версия, сбалансированного дерева.

При вставке нового элемента, в дерево, нужно перестраивать дерево, чтобы оно оставалось сбалансированным.

Как только дерево достигнет 100 элементов, генерируется новое дерево с 5 элементами.
` +

	"```" + `go
package binary

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Key    int
	Height int
	Left   *Node
	Right  *Node
}

type AVLTree struct {
	Root *Node
}

func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

func (t *AVLTree) ToMermaid() string {

}

func height(node *Node) int {

}

func max(a, b int) int {

}

func updateHeight(node *Node) {

}

func getBalance(node *Node) int {

}

func leftRotate(x *Node) *Node {

}

func rightRotate(y *Node) *Node {

}

func insert(node *Node, key int) *Node {

}

func GenerateTree(count int) *AVLTree {

}
` +
	"```" + `

Не обязательно использовать выше описанный код, можно использовать любую реализацию, выдающую сбалансированное бинарное дерево.

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Вывод:

{{< columns >}}
` +
	"```" + `tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph TD
`

const content5 = `
{{/*/* /mermaid */*/}}
` + "```" + `
{{< /columns >}}

{{< mermaid >}}
graph TD
`

const content6 = `
---
menu:
    after:
        name: graph
        weight: 1
title: Построение графа
---

# Построение графа

Нужно написать воркер, который будет строить граф на текущей странице, каждые 5 секунд
От 5 до 30 элементов, случайным образом. Все ноды графа должны быть связаны.
` +
	"```" + `go
type Node struct {
    ID int
    Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
    Links []*Node
}
` + "```" + `

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Граф

{{< columns >}}
` + "```" + `tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
`

const content7 = `
{{</*/* /mermaid */*/>}}
` + "```" + `

<--->

{{< mermaid >}}
graph LR
`
const content8 = `
{{< /mermaid >}}

{{< /columns >}}
`
