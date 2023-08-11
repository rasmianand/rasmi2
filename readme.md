docker run -p 8080:8080 mux-web-service

 The web service will be accessible at `http://localhost:8080`.

## API Endpoints

### Encrypt

Encrypt a value.

- **Endpoint**: `/api/encrypt`
- **Method**: `POST`
- **Request Payload**:

```json
{
 "value": "string_to_encrypt"
}
