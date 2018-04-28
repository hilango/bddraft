package bddraft

/*

//
参考：http://static.xiaorui.cc/raft_design.pdf
https://ongardie.github.io/raft-talk-archive/2015/buildstuff/raftscope-replay/


raft算法的实现pdf

一致性算法
多节点环境


{{
client|客户端
server|服务端
node|节点

Leader|领导
Candidate|候选人
Follower|随从

定时器
Term时间片
Term ID
N/2+1
Heartbeats|心跳



}}

选举成Leader需提供TermID 和 LogIndex
Leader 绝对不会删除自己的日志
客户端自己携带ID帮助raft保持幂等性
一条记录提交了，那么它之前的记录一定都是commited.

节点之间的Term和索引一致, 我们就认为数据是一致的
在一个Term里只会有一个Leader
每个Follower只能选一个Leader

currentTerm
服务器最后一次知道的任期号（初始化为 0，持续递增）

voteFor
在当前获得选票的候选人的 Id

log[]
日志条目集( 状态机指令及TermId )

commitIndex
已知最大的索引值

nextIndex[]
每个follower的下一个索引值


Vote RPC
{
Term|候选人的任期号
candidateid|ID
lastLogIndex|候选人的最后日志的索引值
lastLogTerm|候选人最后日志的任期号

Term|当前的任期号, 用于领导人去更新自己
voteGranted|True or False
}


选举
{
node1
node2
node3

最简单选举：
node1 向其它节点发起选举，其它节点一致投票选它；选举成功

简单选举1：
C-1: time=155,Term=2
F-1: Timer=183,Term=3
F-2: Timer=170,Term3

C-发起选举，由于Term=2, 比其它两个节点小，等到"NO"
Condition比Follwer的term id小,不影响“F”定时器在转  !  C 已得知情况, 故意Vote超时, 等他人选举

简单选举2：
C-1: voteGranted=true, term=2
C-2: NO! Term not match
L-1: RequestVote(term=2)
L-2: RequestVote(term=2)

Same term id wait timeout!


Hard election-1
vote for me
都变成一个term id
not term match\
term conflict
not n/2 +1
vote for me

summery election
过程
定时器触发, followers把current_term_id + 1
改变成candidate状态
发送RequestVoteRPC请求

结果
成功选举
别人被选
重新选

Client
Works with leader
Leader return to response when it commits an entry!
Assign uniquqelD to every command ,Leader store latest ID with response.

Client process
Only log entry!

Log Replication
默认心跳为50ms
默认心跳超时为300ms
每次心跳的时候做 Log entry\commit
超过n/2+1就算成功

Log RPC
Term 领导人的任期号
LeaderID 领导人的ID,以便于跟随者重定向请求
pervLogIndex 新的日志条目紧随之前的索引值
entries[] 需要存储当然日志条目 （表示心跳时为空）；一次性发送多个是为了提高效率
LeaderCommit 领导人已经提交的日志的索引值

Term 当前的任期号，用于领导人去更新自己
Success 跟随者包含了匹配上prevLogIndex和preLogTerm 的日志时间为真

log replication-1|日志

Heartbaet & Append Entries
Heartbaet & Append Entries
Only log entry !

log replication - 2
Leader commit !

log replication - 3
Heartbaet & commit
Heartbaet & commit
Follower commit !
}

------------------

常见疑难杂症
if a node reply timeout ？
F_2 如何保持数据一致性  ?  Leader会重试 !

Leader crash
Log entry Ack
Leader在本地commit后, 发给follower commit 之前crash!Hello 还在么？

Follower crash
F_3  crash重新启动后如何平衡数据.
prevLogIndex

Network Partition
正常情况
Heartbaet & commit

网络分区
两个人怎么够法定人数! ! !

Vote Granted
Request Vote

新集群正常
两个人怎么够法定人数! ! !
Heartbeat & Log entry & commit

网络恢复
网络好了后, 开始抢夺LeaderLe_1 term 小于 Le_2 !
Heartbeart & Append Log Entries

一致性
Heartbeat & Log entry & commit

冲突Split brain
如符合法定人数并产生了N条数据 与 新集群怎么保持数据一致性
覆盖 VS 合并 ?
被分区前有些node没有收到commit ?
timer check

预防Split brain
单播制定节点
指定法定人数 , 每次add\reduce都需要更改
加大timeout , retry
统一 client 入口 , But ⋯
监控脑裂情况, 反查各个node的leader是否一致

复杂一致性
Log compress
Snapshot
Last included index : 6
Last included term : 80
state macheie state:

--------------------------


 */
