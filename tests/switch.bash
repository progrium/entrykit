_build_switch_image() {
  docker build -t entrykit:switchtest - >/dev/null << EOF
FROM entrykit
ENTRYPOINT ["/bin/switch","cmd1=echo app foo is running","cmd2=echo app bar is running","--"]
EOF
}

T_switch() {
  _build_switch_image
  result=$(docker run --rm entrykit:switchtest cmd1)
  if [[ "$result" != "app foo is running" ]]; then
    $T_fail "switch should have choosen cmd1"
    return
fi

  result=$(docker run --rm entrykit:switchtest cmd2)
  [[ "$result" == "app bar is running" ]]
}

