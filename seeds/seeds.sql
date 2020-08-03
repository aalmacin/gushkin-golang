INSERT INTO
  public.activities (description, positive, fund_amt, user_id)
VALUES
  ('Took longer to wake up', TRUE, 10000000, 'test'),
  ('Test 1', TRUE, 2000000, 'test'),
  ('Test 2', false, 2000000, 'test'),
  ('A Test!', TRUE, 10000000, 'test'),
  ('New Spanish Book', TRUE, 50000000, 'test'),
  ('Read books', TRUE, 5000000, 'test'),
  ('Lazy', false, 100000000, 'test'),
  ('New Item', TRUE, 5000000, 'test'),
  ('Pomodoro', TRUE, 10000000, 'test'),
  (
    'Missed Activity with Kids',
    false,
    10000000,
    'test'
  );

INSERT INTO
  public.activities (description, positive, fund_amt, user_id)
VALUES
  ('Did not workout', false, 10000000, 'test'),
  (
    'Did not practice music',
    false,
    10000000,
    'test'
  ),
  (
    'Did not make career improvement activity',
    false,
    10000000,
    'test'
  ),
  ('Wasting time', false, 100000000, 'test'),
  ('Missed Pomodoro Goals', false, 10000000, 'test'),
  (
    'Did not stop after pomodoro',
    false,
    5000000,
    'test'
  ),
  ('Drink 1L water', TRUE, 5000000, 'test'),
  ('Pomodoro', TRUE, 5000000, 'test'),
  ('abc', TRUE, 33000000, 'test'),
  ('pangarap', TRUE, 22000000, 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-05 00:55:43.455', 'test'),
  (2, '2020-03-05 01:00:03.099', 'test'),
  (3, '2020-03-05 22:39:59.026', 'test'),
  (4, '2020-03-05 22:47:36.194', 'test'),
  (5, '2020-03-05 22:47:41.095', 'test'),
  (6, '2020-03-05 23:12:07.126', 'test'),
  (7, '2020-03-05 23:27:52.902', 'test'),
  (8, '2020-03-05 23:27:58.941', 'test'),
  (9, '2020-03-06 22:33:36.521', 'test'),
  (10, '2020-03-06 22:34:13.643', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-09 23:01:11.533', 'test'),
  (2, '2020-03-11 01:12:41.804', 'test'),
  (3, '2020-03-14 00:00:00.000', 'test'),
  (4, '2020-03-13 23:00:00.000', 'test'),
  (5, '2020-03-14 18:22:37.810', 'test'),
  (6, '2020-03-14 18:38:22.387', 'test'),
  (7, '2020-03-14 18:54:41.411', 'test'),
  (8, '2020-03-14 19:05:03.419', 'test'),
  (9, '2020-03-14 19:05:14.976', 'test'),
  (10, '2020-03-14 19:05:30.740', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-14 19:05:38.579', 'test'),
  (2, '2020-03-14 19:12:50.561', 'test'),
  (3, '2020-03-14 19:13:59.799', 'test'),
  (4, '2020-03-15 20:11:07.704', 'test'),
  (5, '2020-03-15 20:11:10.445', 'test'),
  (6, '2020-03-15 20:11:12.956', 'test'),
  (7, '2020-03-15 20:11:13.524', 'test'),
  (8, '2020-03-16 02:10:06.980', 'test'),
  (9, '2020-03-16 02:10:43.223', 'test'),
  (10, '2020-03-16 02:36:35.152', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-16 22:31:58.610', 'test'),
  (2, '2020-03-16 22:39:09.458', 'test'),
  (3, '2020-03-16 22:45:51.495', 'test'),
  (4, '2020-03-16 22:49:14.290', 'test'),
  (5, '2020-03-16 02:28:57.000', 'test'),
  (6, '2020-03-18 21:27:32.378', 'test'),
  (7, '2020-03-18 21:32:34.578', 'test'),
  (8, '2020-03-18 21:32:36.481', 'test'),
  (9, '2020-03-18 21:32:38.081', 'test'),
  (10, '2020-03-18 21:32:39.478', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-18 21:32:40.924', 'test'),
  (2, '2020-03-18 21:32:43.439', 'test'),
  (3, '2020-03-18 21:32:44.961', 'test'),
  (4, '2020-03-18 21:32:46.462', 'test'),
  (5, '2020-03-18 21:33:02.419', 'test'),
  (6, '2020-03-19 09:54:42.815', 'test'),
  (7, '2020-03-19 12:31:59.011', 'test'),
  (8, '2020-03-19 13:59:25.217', 'test'),
  (9, '2020-03-19 13:59:28.420', 'test'),
  (10, '2020-03-19 13:59:40.982', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-19 13:59:43.640', 'test'),
  (2, '2020-03-19 17:55:14.530', 'test'),
  (3, '2020-03-19 17:56:15.772', 'test'),
  (4, '2020-03-19 17:56:51.530', 'test'),
  (5, '2020-03-19 17:56:53.533', 'test'),
  (6, '2020-03-19 17:56:55.293', 'test'),
  (7, '2020-03-19 22:58:00.454', 'test'),
  (8, '2020-03-19 22:58:04.596', 'test'),
  (9, '2020-03-20 22:27:45.646', 'test'),
  (10, '2020-03-20 22:27:46.289', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-20 22:27:49.667', 'test'),
  (2, '2020-03-20 22:27:51.966', 'test'),
  (3, '2020-03-20 22:27:53.489', 'test'),
  (4, '2020-03-20 22:27:55.169', 'test'),
  (5, '2020-03-20 22:27:56.671', 'test'),
  (6, '2020-03-20 22:27:58.227', 'test'),
  (7, '2020-03-20 22:27:59.926', 'test'),
  (8, '2020-03-20 22:28:01.831', 'test'),
  (9, '2020-03-20 22:28:08.247', 'test'),
  (10, '2020-03-20 22:28:15.749', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (1, '2020-03-20 22:59:56.098', 'test'),
  (2, '2020-03-20 23:02:02.479', 'test'),
  (3, '2020-03-21 22:31:09.508', 'test'),
  (4, '2020-03-21 22:31:11.847', 'test'),
  (5, '2020-03-21 23:29:08.147', 'test'),
  (6, '2020-03-21 23:29:49.926', 'test'),
  (7, '2020-03-21 23:30:30.185', 'test'),
  (8, '2020-03-21 23:30:44.932', 'test'),
  (9, '2020-03-21 23:30:45.750', 'test'),
  (10, '2020-03-21 23:30:46.507', 'test');

INSERT INTO
  actions (activity_id, action_timestamp, user_id)
VALUES
  (4, '2020-03-21 23:31:44.226', 'test'),
  (2, '2020-03-21 23:31:46.467', 'test'),
  (8, '2020-03-21 23:32:03.586', 'test'),
  (9, '2020-03-21 23:31:44.226', 'test'),
  (3, '2020-03-21 23:31:46.467', 'test'),
  (9, '2020-03-21 23:32:03.586', 'test');

INSERT INTO
  public.wishes (
    description,
    price,
    "source",
    priority,
    STATUS,
    user_id
  )
VALUES
  (
    'Another new',
    33,
    'dfsdfsf',
    'MEDIUM',
    'bought',
    'test'
  ),
  (
    'Raizza',
    323,
    'Here man',
    'LOW',
    'disabled',
    'test'
  ),
  (
    'ajedrez',
    334,
    'Aldrin',
    'MEDIUM',
    'bought',
    'test'
  ),
  (
    'new udemy course',
    15,
    'udemy.com',
    'MEDIUM',
    'disabled',
    'test'
  ),
  (
    'New Book',
    1000000,
    'www.amazon.ca',
    'VERY_LOW',
    'bought',
    'test'
  ),
  (
    'New watch band',
    10000000,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'New watch band',
    10000000,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'New watch band',
    10000000,
    NULL,
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'New watch band',
    10000000,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'book11',
    11000000,
    'Trenta',
    'VERY_HIGH',
    'bought',
    'test'
  );

INSERT INTO
  public.wishes (
    description,
    price,
    "source",
    priority,
    STATUS,
    user_id
  )
VALUES
  ('New', 22, 'Here', 'VERY_LOW', 'bought', 'test'),
  (
    'book3',
    34,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  ('book2', 2, 'NULL', 'VERY_HIGH', 'bought', 'test'),
  (
    'refresh now',
    33,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'book99',
    990,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'New book for spanish learning',
    344000000,
    'NULL',
    'VERY_HIGH',
    'not_bought',
    'test'
  ),
  (
    'Madden 20',
    80000000,
    'NULL',
    'VERY_HIGH',
    'not_bought',
    'test'
  ),
  (
    'Teddy bear',
    10000000,
    'NULL',
    'VERY_HIGH',
    'bought',
    'test'
  ),
  (
    'book114',
    11400000,
    'Trenta',
    'VERY_HIGH',
    'bought',
    'test'
  );