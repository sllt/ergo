package ergo

import (
	"github.com/sllt/ergo/etf"
	"github.com/sllt/ergo/gen"
	"github.com/sllt/ergo/node"
	"time"
)

var (
	DefaultNode node.Node
)

func init() {
	DefaultNode, _ = StartNode("ergo@localhost", "ergo", node.Options{})
}

// Name returns node name
func Name() string {
	return DefaultNode.Name()
}

// IsAlive returns true if node is still alive
func IsAlive() bool {
	return DefaultNode.IsAlive()
}

// Uptime returns node uptime in seconds
func Uptime() int64 {
	return DefaultNode.Uptime()
}

// ListEnv returns a map of configured Node environment variables.
func ListEnv() map[gen.EnvKey]interface{} {
	return DefaultNode.ListEnv()
}

// SetEnv set node environment variable with given name. Use nil value to remove variable with given name. Ignores names with "ergo:" as a prefix.
func SetEnv(name gen.EnvKey, value interface{}) {
	DefaultNode.SetEnv(name, value)
}

// Env returns value associated with given environment name.
func Env(name gen.EnvKey) interface{} {
	return DefaultNode.Env(name)
}

func RegisterName(name string, pid etf.Pid) error {
	return DefaultNode.RegisterName(name, pid)
}
func UnregisterName(name string) error {
	return DefaultNode.UnregisterName(name)
}

func LoadedApplications() []gen.ApplicationInfo {
	return DefaultNode.LoadedApplications()
}
func WhichApplications() []gen.ApplicationInfo {
	return DefaultNode.WhichApplications()
}
func ApplicationInfo(name string) (gen.ApplicationInfo, error) {
	return DefaultNode.ApplicationInfo(name)
}
func ApplicationLoad(app gen.ApplicationBehavior, args ...etf.Term) (string, error) {
	return DefaultNode.ApplicationLoad(app, args...)
}
func ApplicationUnload(appName string) error {
	return DefaultNode.ApplicationUnload(appName)
}
func ApplicationStart(appName string, args ...etf.Term) (gen.Process, error) {
	return DefaultNode.ApplicationStart(appName, args...)
}
func ApplicationStartPermanent(appName string, args ...etf.Term) (gen.Process, error) {
	return DefaultNode.ApplicationStartPermanent(appName, args...)
}
func ApplicationStartTransient(appName string, args ...etf.Term) (gen.Process, error) {
	return DefaultNode.ApplicationStartTransient(appName, args...)
}
func ApplicationStop(appName string) error {
	return DefaultNode.ApplicationStop(appName)
}

func ProvideRemoteSpawn(name string, object gen.ProcessBehavior) error {
	return DefaultNode.ProvideRemoteSpawn(name, object)
}
func RevokeRemoteSpawn(name string) error {
	return DefaultNode.RevokeRemoteSpawn(name)
}

// AddStaticRoute adds static route for the given name
func AddStaticRoute(node string, host string, port uint16, options node.RouteOptions) error {
	return DefaultNode.AddStaticRoute(node, host, port, options)
}

// AddStaticRoutePort adds static route for the given node name which makes node skip resolving port process
func AddStaticRoutePort(node string, port uint16, options node.RouteOptions) error {
	return DefaultNode.AddStaticRoutePort(node, port, options)
}

// AddStaticRouteOptions adds static route options for the given node name which does regular port resolving but applies static options
func AddStaticRouteOptions(node string, options node.RouteOptions) error {
	return DefaultNode.AddStaticRouteOptions(node, options)
}

// Remove static route removes static route with given name
func RemoveStaticRoute(name string) bool {
	return DefaultNode.RemoveStaticRoute(name)
}

// StaticRoutes returns list of routes added using AddStaticRoute
func StaticRoutes() []node.Route {
	return DefaultNode.StaticRoutes()
}

// StaticRoute returns Route for the given name. Returns false if it doesn't exist.
func StaticRoute(name string) (node.Route, bool) {
	return DefaultNode.StaticRoute(name)
}

func AddProxyRoute(proxy node.ProxyRoute) error {
	return DefaultNode.AddProxyRoute(proxy)
}

func RemoveProxyRoute(name string) bool {
	return DefaultNode.RemoveProxyRoute(name)
}

// ProxyRoutes returns list of proxy routes added using AddProxyRoute
func ProxyRoutes() []node.ProxyRoute {
	return DefaultNode.ProxyRoutes()
}

// ProxyRoute returns proxy route added using AddProxyRoute
func ProxyRoute(name string) (node.ProxyRoute, bool) {
	return DefaultNode.ProxyRoute(name)
}

// Resolve
func Resolve(node string) (node.Route, error) {
	return DefaultNode.Resolve(node)
}

// ResolveProxy resolves proxy route. Checks for the proxy route added using AddProxyRoute.
// If it wasn't found makes request to the registrar.
func ResolveProxy(node string) (node.ProxyRoute, error) {
	return DefaultNode.ResolveProxy(node)
}

// Returns Registrar interface
func Registrar() node.Registrar {
	return DefaultNode.Registrar()
}

// Connect sets up a connection to node
func Connect(node string) error {
	return DefaultNode.Connect(node)
}

// Disconnect close connection to the node
func Disconnect(node string) error {
	return DefaultNode.Disconnect(node)
}

// Nodes returns the list of connected nodes
func Nodes() []string {
	return DefaultNode.Nodes()
}

// NodesIndirect returns the list of nodes connected via proxies
func NodesIndirect() []string {
	return DefaultNode.NodesIndirect()
}

// NetworkStats returns network statistics of the connection with the node. Returns error
// ErrUnknown if connection with given node is not established.
func NetworkStats(name string) (node.NetworkStats, error) {
	return DefaultNode.NetworkStats(name)
}

func Links(process etf.Pid) []etf.Pid {
	return DefaultNode.Links(process)
}
func Monitors(process etf.Pid) []etf.Pid {
	return DefaultNode.Monitors(process)
}
func MonitorsByName(process etf.Pid) []gen.ProcessID {
	return DefaultNode.MonitorsByName(process)
}
func MonitoredBy(process etf.Pid) []etf.Pid {
	return DefaultNode.MonitoredBy(process)
}

func Stats() node.NodeStats {
	return DefaultNode.Stats()
}

func Stop() {
	DefaultNode.Stop()
}
func Wait() {
	DefaultNode.Wait()
}
func WaitWithTimeout(d time.Duration) error {
	return DefaultNode.WaitWithTimeout(d)
}

// Spawn spawns a new process
func Spawn(name string, opts gen.ProcessOptions, object gen.ProcessBehavior, args ...etf.Term) (gen.Process, error) {
	return DefaultNode.Spawn(name, opts, object, args...)
}

// ProcessByName returns Process for the given name.
// Returns nil if it doesn't exist (not found) or terminated.
func ProcessByName(name string) gen.Process {
	return DefaultNode.ProcessByName(name)
}

// ProcessByPid returns Process for the given Pid.
// Returns nil if it doesn't exist (not found) or terminated.
func ProcessByPid(pid etf.Pid) gen.Process {
	return DefaultNode.ProcessByPid(pid)
}

// ProcessByAlias returns Process for the given alias.
// Returns nil if it doesn't exist (not found) or terminated
func ProcessByAlias(alias etf.Alias) gen.Process {
	return DefaultNode.ProcessByAlias(alias)
}

// ProcessInfo returns the details about given Pid
func ProcessInfo(pid etf.Pid) (gen.ProcessInfo, error) {
	return DefaultNode.ProcessInfo(pid)
}

// ProcessList returns the list of running processes
func ProcessList() []gen.Process {
	return DefaultNode.ProcessList()
}

// MakeRef creates an unique reference within this node
func MakeRef() etf.Ref {
	return DefaultNode.MakeRef()
}

// IsAlias checks whether the given alias is belongs to the alive process on this node.
// If the process died all aliases are cleaned up and this function returns
// false for the given alias. For alias from the remote node always returns false.
func IsAlias(alias etf.Alias) bool {
	return DefaultNode.IsAlias(alias)
}

// IsMonitor returns true if the given references is a monitor
func IsMonitor(ref etf.Ref) bool {
	return DefaultNode.IsMonitor(ref)
}

// RegisterBehavior
func RegisterBehavior(group, name string, behavior gen.ProcessBehavior, data interface{}) error {
	return DefaultNode.RegisterBehavior(group, name, behavior, data)
}

// RegisteredBehavior
func RegisteredBehavior(group, name string) (gen.RegisteredBehavior, error) {
	return DefaultNode.RegisteredBehavior(group, name)
}

// RegisteredBehaviorGroup
func RegisteredBehaviorGroup(group string) []gen.RegisteredBehavior {
	return DefaultNode.RegisteredBehaviorGroup(group)
}

// UnregisterBehavior
func UnregisterBehavior(group, name string) error {
	return DefaultNode.UnregisterBehavior(group, name)
}
