<template>
	<view class="content">
		<image class="logo" src="/static/logo.png"></image>
		<view class="text-area">
			<text class="title">{{title}}</text>
		</view>
		<view>
			<text>{{textvalue}}</text><!-- 这里演示了组件值的绑定 -->
			<button :type="buttontype" @click="changetextvalue()">修改为789</button><!--这里演示了属性和事件的绑定-->
		</view>
		<!--插值-->
		<view>
			<view>Message：{{msg}}</view> <!--文本-->
			<view>{{ok ? 'YES' : 'NO'}}</view> <!--javascript表达式-->
			<view>{{message.split('').reverse().join('')}}</view><!--javascript表达式 拆解，反向，拼接-->
			<view>socket接收的数据：{{sdata}}</view>
		</view>
	</view>
</template>

<script>
	var globalvar = 1 // 整个页面都可以被使用
	
	
	export default {
		data() {
			return {
				title: 'Hello',
				textvalue:"123",
				buttontype:"primary",
				msg:"文本插值",
				ok:true,
				message:'hello Vue!',
				sdata:'--'
			}
		},
		// 这个是小程序的
		onLoad() {
		 	this.textvalue="456"; // 这里修改textvalue的值
			// uni.request({
			//     url: 'http://127.0.0.1:8080/', //仅为示例，并非真实接口地址。
			//     // data: {
			//     //     text: 'uni.request'
			//     // },
			//     // header: {
			//     //     'custom-header': 'hello' //自定义请求头信息
			//     // },
			//     success: (res) => {
			//         console.log(res.data);
			//         this.title = res.data;
			//     },
			// });
		},
		// 这个是vue的生命周期
		mounted(){
			uni.connectSocket({
			  url: 'ws://127.0.0.1:8080/socket?user=1'
			});
			
			uni.onSocketOpen(function (res) {
			  console.log('WebSocket连接已打开！' + res.data);
			  uni.sendSocketMessage({
			        data: globalvar
			  });
			});
			
			uni.onSocketError(function (res) {
			  console.log('WebSocket连接打开失败，请检查！' + res.data);
			});
			
			// uni.onSocketMessage(function (res) {
			//   console.log('收到服务器内容：' + res.data);
			  
			//   globalvar = parseInt(res.data) +1 ;
			  
			//   _this.sdata = globalvar
			  
			//   setTimeout(SendMsg,1000);
			// });
			
			uni.onSocketMessage( (res) => {
			  console.log('收到服务器内容：' + res.data);
			  
			  globalvar = parseInt(res.data) +1 ;
			  
			  this.sdata = globalvar
			  
			  setTimeout(this.SendMsg,1000);
			 
			 // uni.sendSocketMessage({
			 //       data: globalvar
			 // });
			  
			});
		},
		methods: {
			changetextvalue(){
				const _this = this;
				this.textvalue= parseInt(this.textvalue) +1; // 这里修改textvalue的值
				
				// uni.request({
				// 	url:'http://127.0.0.1:8080/',
				// 	success:(res) => {
				// 		 console.log(res.data);
				// 		this.textvalue  = res.data;
				// 	}
				// });
				
				
	
			
			},
			 SendMsg(){
				uni.sendSocketMessage({
				      data: globalvar
				});
			}
		}
	}
</script>

<style>
	.content {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
	}

	.logo {
		height: 200rpx;
		width: 200rpx;
		margin-top: 200rpx;
		margin-left: auto;
		margin-right: auto;
		margin-bottom: 50rpx;
	}

	.text-area {
		display: flex;
		justify-content: center;
	}

	.title {
		font-size: 36rpx;
		color: #8f8f94;
	}
</style>
