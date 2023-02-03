<template>
	<view class="content">
		<view class="dorm-box">
			<view class="out-title">猪舍监控</view>
			<view class="items-box">
				<u-empty textSize="28" v-if="houses.length === 0" mode="data" text="无猪舍记录"></u-empty>
				<view class="my-item" @click="toHouseMonitor(item.id)" v-for="(item,index) in houses" :key="index">
					<view class="in-title">
						{{item.houseName}}
					</view>
					<view class="dorm-monitor">
						<view class="dorm-img">
							<image src="/static/image/temperature.png" mode=""></image>
						</view>
						<text class="dorm-data">{{ item.temperature }}</text>
						<view class="dorm-img">
							<image src="/static/image/wifi-off.png" mode=""></image>
						</view>
						<text class="dorm-data">{{ item.outlineNum }}台</text>
						<view class="dorm-img">
							<image src="/static/image/warn.png" mode=""></image>
						</view>
						<text class="dorm-data">{{ item.alarmNum }}次</text>
					</view>
				</view>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				houses: []
			}
		},
		mounted() {},
		onPullDownRefresh() {
			this.loadFarmInfo(() => {
				uni.stopPullDownRefresh();
			});
		},
		onShow() {
			this.loadFarmInfo(() => {});
		},
		methods: {
			async loadFarmInfo(callback) {
				const farmHouseData = await this.$http('farm.houses', {
					farmId: farmData.data.farmId
				});
				if (farmHouseData && farmHouseData.code == 200) {
					this.houses = farmHouseData.data.houses
				}
				callback && callback();
			},
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
	
</style>
