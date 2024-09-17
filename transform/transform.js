/*  1. While you can write multiple functions, the main
   function called for your transformation is the transform function.

2. The only argument acceptable in the transform function is the
 payload data.

3. The transform method must return a value.

4. Console logs lust be written like this:
console.log('%j', logged_item) to get printed in the log below.

5. The output payload from the function should be in this format
    {
        "owner_id": "string, optional",
        "event_type": "string, required",
        "data": "object, required",
        "custom_headers": "object, optional",
        "idempotency_key": "string, optional"
        "endpoint_id": "string, depends",
    }

6. The endpoint_id field is only required when sending event to
a single endpoint. */

function transform(payload) {
    // Transform function here
    return {
        "owner_id": payload.business_id,
        "event_type": payload.event_type,
        "data": {
            "event_type": payload.event_type,
            "data": payload
        },
        "idempotency_key": payload.id,
        "custom_headers": {
            "x-convoy-message-type": "fanout"
        }
    }
}
