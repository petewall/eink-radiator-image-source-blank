# eInk Radiator Image Source: Blank

![CI](https://ci.petewall.net/api/v1/teams/main/pipelines/eink-radiator/jobs/test-image-source-blank/badge)

Generates an image with a single color.

```bash
blank generate --config config.json --height 300 --width 400
```

## Configuration

The only configuration is the image color. The config file can be in JSON or YAML format:

| field | default | required | description |
|-------|---------|----------|-------------|
| color |         | Yes      | The color   |

## Example

```yaml
---
color: orange
```

![An image rendering the orange example](test/outputs/orange.png)
