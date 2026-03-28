-- ============================================================
--  База данных: Учёт инструментов и материалов
--  3 таблицы: employees, instruments, issues
-- ============================================================

-- Таблица 1: Сотрудники
CREATE TABLE IF NOT EXISTS employees (
    id         SERIAL PRIMARY KEY,
    full_name  VARCHAR(100) NOT NULL,
    position   VARCHAR(100) NOT NULL,
    department VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Таблица 2: Инструменты
CREATE TABLE IF NOT EXISTS instruments (
    id               SERIAL PRIMARY KEY,
    name             VARCHAR(200) NOT NULL,
    inventory_number VARCHAR(50)  NOT NULL UNIQUE,
    category         VARCHAR(100) NOT NULL,
    status           VARCHAR(30)  NOT NULL DEFAULT 'in_stock',
    description      TEXT,
    created_at       TIMESTAMP DEFAULT NOW()
);

-- Таблица 3: Выдачи (связывает instruments и employees)
CREATE TABLE IF NOT EXISTS issues (
    id                   SERIAL PRIMARY KEY,
    instrument_id        INT NOT NULL REFERENCES instruments(id) ON DELETE CASCADE,
    employee_id          INT NOT NULL REFERENCES employees(id)   ON DELETE CASCADE,
    issue_date           DATE NOT NULL DEFAULT CURRENT_DATE,
    expected_return_date DATE,
    return_date          DATE,
    note                 TEXT,
    created_at           TIMESTAMP DEFAULT NOW()
);

-- Тестовые данные
INSERT INTO employees (full_name, position, department) VALUES
    ('Иванов Иван Иванович',   'Слесарь',    'Производственный цех'),
    ('Петров Пётр Петрович',   'Электрик',   'Технический отдел'),
    ('Сидоров Алексей Юрьевич','Сварщик',    'Производственный цех');

INSERT INTO instruments (name, inventory_number, category, status, description) VALUES
    ('Дрель ударная Bosch',   'INV-001', 'Электроинструмент', 'in_stock',  'Мощность 800Вт'),
    ('Молоток слесарный',      'INV-002', 'Ручной инструмент', 'in_stock',  'Вес 0.5 кг'),
    ('Шуруповёрт Makita',      'INV-003', 'Электроинструмент', 'in_stock',  'Аккумуляторный 18В'),
    ('Набор отвёрток 8шт',     'INV-004', 'Ручной инструмент', 'in_stock',  'Phillips и Slotted'),
    ('Болгарка 125мм',         'INV-005', 'Электроинструмент', 'in_stock',  'УШМ 1200Вт');
