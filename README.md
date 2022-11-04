# eInk Radiator Image Source: Blank

![CI](https://ci.petewall.net/api/v1/teams/main/pipelines/eink-radiator/jobs/test-image-source-blank/badge)

Generates an image with a single color.

```bash
blank --config config.json --height 300 --width 400
```

## Configuration

The only configuration is the image color. The config file can be in JSON or YAML format:

```json
{
    "color": "purple"
}
```
