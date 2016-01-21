_render_in_docker() {
  declare templateFile=${1:?required}
  docker run --rm \
      -v $PWD/${templateFile}:/work/template.tmpl \
      -w /work \
      -e NAME=megalith \
      --entrypoint render \
      entrykit \
        template -- cat template
}

T_render() {
  result=$(_render_in_docker tests/hello.tmpl)
  [[ "$result" == "hello world" ]]
}

T_render_posix_var() {
  result=$(_render_in_docker tests/posix.tmpl)
  [[ "$result" == "hello megalith" ]]
}
