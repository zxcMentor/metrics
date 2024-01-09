
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
```go
type Node struct {
    ID int
    Name string
	Form string // "circle", "rect", "square", "ellipse", "round-rect", "rhombus"
    Links []*Node
}
```

## Mermaid Chart

[MermaidJS](https://mermaid-js.github.io/) is library for generating svg charts and diagrams from text.

## Граф

{{< columns >}}
```tpl
{{</*/* mermaid [class="text-center"]*/*/>}}
graph LR
A[A] --> B[B]
A --> D[D]
B --> E[E]
B --> C[C]
C --> D
C --> E
D --> E
D --> F([F])
E --> F
E --> G[G]
F --> G
F --> H((H))
G --> H
H --> I[I]
H --> J[J]
I --> J

{{</*/* /mermaid */*/>}}
```

<--->

{{< mermaid >}}
graph LR
A[A] --> B[B]
A --> D[D]
B --> E[E]
B --> C[C]
C --> D
C --> E
D --> E
D --> F([F])
E --> F
E --> G[G]
F --> G
F --> H((H))
G --> H
H --> I[I]
H --> J[J]
I --> J

{{< /mermaid >}}

{{< /columns >}}
