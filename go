Test:
  go test -v ./... -args --logtostderr -v=10   go test -v ./... -args --logtostderr -v=10
  go test kubectlfzf/pkg/resourcewatcher --args --logtostderr -v=10

Compile:
  go test kubectlfzf/pkg/resourcewatcher -c
  dlv test kubectlfzf/pkg/resourcewatcher
