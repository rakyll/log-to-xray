# logtoxray

Write to logs, get X-Ray traces. No distributing
tracing instrumenation library required.

logtoxray processes aggregates log entries to
create X-Ray segments.

## Usage

Users should include trace_id and span_id in every entry.

In order to start a new segment, log an entry
with start_time.

```
{
    "trace_id": "...",
    "span_id": "...",
    "name": "auth.CurrentUser",
    "start_time": "2021-06-28 17:09:12.0 -0700 PDT"
}
```

Append as many as changes with new log entries.
Note that trace_id, span_id, start_time and end_time
are immutable.

```
{
    "trace_id": "...",
    "span_id": "...",
    "attrs": {
        "service": "auth",
        "region": "us-east-1"
    }
}
```

In order to finish a segment, write an entry
with end_time:

```
{
    "trace_id": "...",
    "span_id": "...",
    "end_time": "2021-06-28 17:10:26.086625 -0700 PDT"
}
```
