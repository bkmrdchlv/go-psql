INSERT INTO something (int, string, bool, float, map, struct, pointer, interface, timestamp)
VALUES (
  1233,  -- Remove the extra space before the number
  'This is a string',
  true,
  3,
  '{"key1": "value1", "key2": 2}',
  '{"field1": "data1", "field2": 1.23}',
  'pointer_value',
  '{"method": "function", "params": []}',
  1687366400 -- Timestamp for August 25, 2024
);