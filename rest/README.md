# chat.*

```js
/api/v1/chat.delete
Deletes an existing chat message.
Link
/api/v1/chat.followMessage
Follows an existing chat message.
Link
/api/v1/chat.getDeletedMessages
Retrieves the deleted messages since specific date.
Link
/api/v1/chat.getDiscussions
Retrieves the discussions of a room.
Link
/api/v1/chat.getMentionedMessages
Retrieves the mentioned messages.
Link
/api/v1/chat.getMessage
Retrieves a single chat message.
Link
/api/v1/chat.getMessageReadReceipts (可能被弃用)
Retrieves message read receipts.
Link
/api/v1/chat.getPinnedMessages
Retrieve pinned messages from a room.
Link
/api/v1/chat.getSnippetedMessages (暂时没写 模板信息 可能需要管理员权限)
Retrieves snippeted messages.
Link
/api/v1/chat.getSnippetedMessageById (暂时没写 模板信息 可能需要管理员权限)
Retrieves snippeted message by id.
Link
/api/v1/chat.getStarredMessages
Retrieves the starred messages.
Link
/api/v1/chat.getThreadMessages
Retrieves thread's messages.
Link
/api/v1/chat.getThreadsList
Retrieves channel's threads.
Link
/api/v1/chat.ignoreUser
Ignores an user from a chat.
Link
/api/v1/chat.pinMessage
Pins a chat message to the message's channel.
Link
/api/v1/chat.postMessage
Posts a new chat message.
Link
/api/v1/chat.react
Sets/unsets the user's reaction to an existing chat message.
Link
/api/v1/chat.reportMessage  (作用不明 未测试)
Reports a message.
Link
/api/v1/chat.search
Search for messages in a channel.
Link
/api/v1/chat.starMessage
Stars a chat message for the authenticated user.
Link
/api/v1/chat.sendMessage  (有Bug (去掉返回信息的avator alias 就行))
Send new chat message.
Link
/api/v1/chat.syncThreadMessages
Retrieves synced thread's messages.
Link
/api/v1/chat.syncThreadsList
Retrieves thread's synced channel threads.
Link
/api/v1/chat.unfollowMessage
Unfollows an existing chat message.
Link
/api/v1/chat.unPinMessage
Removes the pinned status of the provided chat message.
Link
/api/v1/chat.unStarMessage
Removes the star on the chat message for the authenticated user.
Link
/api/v1/chat.update
Updates the text of the chat message.
Link
```

# Channel.*

```js
/api/v1/channels.addAll  (false Access to Method Forbidden [403] 可能需要管理员权限)
Adds all of the users on the server to a channel.
Link
/api/v1/channels.addLeader
Gives the role of Leader for a user in the current channel.
Link
/api/v1/channels.addModerator
Gives the role of moderator to a user in a channel.
Link
/api/v1/channels.addOwner
Gives the role of owner to a user in a channel.
Link
/api/v1/channels.anonymousread  (暂时没写 需要设置匿名聊天)
Gets the messages in public channels to an anonymous user
Link
/api/v1/channels.archive
Archives a channel.
Link
/api/v1/channels.close
Removes a channel from a user's list of channels.
Link
/api/v1/channels.counters
Gets channel counters.
Link
/api/v1/channels.create
Creates a new channel.
Link
/api/v1/channels.delete
Removes a channel.
Link
/api/v1/channels.getAllUserMentionsByChannel
Get all the mentions of a channel.
Link
/api/v1/channels.files
Gets a list of files from a channel.
Link
/api/v1/channels.getIntegrations  (暂时没写 需要开启integrations功能)
Gets the channel's integration.
Link
/api/v1/channels.history
Retrieves the messages from a channel.
Link
/api/v1/channels.info
Gets a channel's information.
Link
/api/v1/channels.invite
Adds a user to a channel.
Link
/api/v1/channels.join
Joins yourself to a channel.
Link
/api/v1/channels.kick
Removes a user from a channel.
Link
/api/v1/channels.leave
Removes the calling user from a channel.
Link
/api/v1/channels.list
Retrieves all of the channels from the server.
Link
/api/v1/channels.list.joined
Gets only the channels the calling user has joined.
Link
/api/v1/channels.members
Retrieves all channel users.
Link
/api/v1/channels.messages  (返回结构可能不完整）
Retrieves all channel messages.
Link
/api/v1/channels.moderators
List all moderators of a channel.
Link
/api/v1/channels.online
List all online users of a channel.
Link
/api/v1/channels.open
Adds the channel back to the user's list of channels.
Link
/api/v1/channels.removeleader
Removes the role of Leader for a user in the current channel.
Link
/api/v1/channels.removeModerator
Removes the role of moderator from a user in a channel.
Link
/api/v1/channels.removeOwner
Removes the role of owner from a user in a channel.
Link
/api/v1/channels.rename
Changes a channel's name.
Link
/api/v1/channels.roles
Gets the user's roles in the channel.
Link
/api/v1/channels.setAnnouncement
Sets a channel's announcement.
Link
/api/v1/channels.setCustomFields (测试通过  作用机制不明确)
Sets a channel's custom fields.
Link
/api/v1/channels.setDefault
Sets a channel's default status.
Link
/api/v1/channels.setDescription
Sets a channel's description.
Link
/api/v1/channels.setJoinCode
Sets the channel's code required to join it.
Link
/api/v1/channels.setPurpose  (没写 文档里这么说 Sets the description for the channel (the same as channels.setDescription, obsolete).)
Sets a channel's description.
Link
/api/v1/channels.setReadOnly
Sets whether a channel is read only or not.
Link
/api/v1/channels.setTopic
Sets a channel's topic.
Link
/api/v1/channels.setType
Sets the type of room the channel should be.
Link
/api/v1/channels.unarchive (需要管理员权限)
Unarchives a channel.
Links
```

# Group.*
```js
/api/v1/groups.addAll (需要管理员权限)
Adds all of the users on the server to a private group.
Link
/api/v1/groups.addLeader
Gives the role of Leader for a user in the current group.
Link
/api/v1/groups.addModerator
Gives the role of moderator to a user in a group.
Link
/api/v1/groups.addOwner
Gives the role of owner to a user in a group.
Link
/api/v1/groups.archive
Archives a private group.
Link
/api/v1/groups.close
Removes a private group from the list of groups.
Link
/api/v1/groups.counters
Gets group counters.
Link
/api/v1/groups.create
Creates a new private group.
Link
/api/v1/groups.delete
Removes a private group.
Link
/api/v1/groups.files
Gets a list of files from a private group.
Link
/api/v1/groups.getIntegrations (暂时没写 需要开启相关功能)
Gets the integrations assigned to the group.
Link
/api/v1/groups.history
Retrieves the messages from a private group.
Link
/api/v1/groups.info
Gets the information about a private group.
Link
/api/v1/groups.invite
Adds a user to the private group.
Link
/api/v1/groups.kick
Removes a user from a private group.
Link
/api/v1/groups.leave
Removes the calling user from the private group.
Link
/api/v1/groups.list
List the private groups the caller is part of.
Link
/api/v1/groups.listAll
List all the private groups.
Link
/api/v1/groups.moderators
List all moderators of a group.
Link
/api/v1/groups.members
Gets the users of participants of a private group.
Link
/api/v1/groups.messages
Retrieves all group messages.
Link
/api/v1/groups.online
List all online users of a group.
Link
/api/v1/groups.open
Adds the private group back to the list of groups.
Link
/api/v1/groups.removeLeader
Removes the role of Leader for a user in the current group.
Link
/api/v1/groups.removeModerator
Removes the role of moderator from a user in a group.
Link
/api/v1/groups.removeOwner
Removes the role of owner from a user in a group.
Link
/api/v1/groups.rename
Changes the name of the private group.
Link
/api/v1/groups.roles
Gets the user's roles in the private group.
Link
/api/v1/groups.setAnnouncement
Sets a group's announcement.
Link
/api/v1/groups.setCustomFields (不知道干啥的)
Sets private group's custom fields.
Link
/api/v1/groups.setDescription
Sets a private group's description.
Link
/api/v1/groups.setPurpose (没写 功能跟上一个重复)
Sets a private group's description.
Link
/api/v1/groups.setReadOnly
Sets whether the room is read only or not.
Link
/api/v1/groups.setTopic
Sets a private group's topic.
Link
/api/v1/groups.setType
Sets the type of room this group will be.
Link
/api/v1/groups.unarchive (需要管理员权限)
Unarchives a private group.
Link
```
