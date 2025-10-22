components {
  id: "block"
  component: "/main/scripts/block.script"
}
embedded_components {
  id: "bg"
  type: "sprite"
  data: "default_animation: \"block_bg\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/main.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "el"
  type: "sprite"
  data: "default_animation: \"block_1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/main.atlas\"\n"
  "}\n"
  ""
}
