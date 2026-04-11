# 💳 Payment Gateway - High-Performance Financial System

Este projeto é um **Gateway de Pagamento** de alta performance desenvolvido em **Go**, projetado para garantir segurança financeira, escalabilidade e integridade de dados em ambientes de alto volume de transações.

O foco principal é resolver desafios críticos do setor de Fintech, como a prevenção de pagamentos duplicados e a garantia de que nenhuma transação seja perdida entre serviços.

---

## 🚀 Tecnologias e Conceitos Implementados

* **Go (Golang):** Escolhido pela alta performance e gerenciamento eficiente de concorrência via *goroutines*.
* **PostgreSQL:** Banco de dados relacional para garantir transações ACID e integridade dos dados financeiros.
* **Redis (Idempotency):** Implementação de chaves de idempotência para evitar cobranças duplicadas em caso de instabilidade de rede ou cliques duplos.
* **Apache Kafka & Outbox Pattern:** Arquitetura orientada a eventos para garantir que a comunicação entre serviços seja resiliente e que nenhum evento financeiro seja perdido.
* **Double-Entry Bookkeeping:** Sistema de partidas dobradas para auditoria e rastreabilidade total de saldos.

---

## 🏗️ Arquitetura e Resiliência

Para atingir o nível de confiança exigido por sistemas bancários, o gateway utiliza padrões arquiteturais avançados:

### 🛡️ Prevenção de Duplicidade (Idempotency)
Cada requisição de pagamento exige um cabeçalho `Idempotency-Key`. O sistema utiliza o **Redis** para validar se aquela transação já foi processada nos últimos 30 minutos, retornando o resultado em cache caso a chave se repita.

### 📦 Entrega Garantida (Outbox Pattern)
Para evitar inconsistências entre o banco de dados e o broker de mensagens, implementamos o **Outbox Pattern**. A transação e o evento de domínio são gravados na mesma transação atômica do PostgreSQL, garantindo que o Kafka seja notificado apenas se os dados forem persistidos com sucesso.



---

## 🛠️ Como rodar o projeto

### Pré-requisitos
* Docker & Docker Compose
* Go 1.21+

### Passo a passo
1. Clone o repositório:
   ```bash
   git clone [https://github.com/seu-usuario/payment-gateway.git](https://github.com/seu-usuario/payment-gateway.git)
